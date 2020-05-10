package queue

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/delivery"
	"github.com/tamaApotek/tama-go-server/domain/apperror"
)

type handler struct {
	queueUsecase Usecase
}

func NewHandler(r *gin.RouterGroup, d delivery.Delivery, queueUsecase Usecase) {
	h := &handler{queueUsecase}

	r.POST("/", h.Add)
}

func handleError(c *gin.Context, err error) {
	r := delivery.Response{}

	switch {
	case
		errors.Is(err, ErrInvalidDate),
		errors.Is(err, ErrInvalidDoctor):

		wrapped := errors.Unwrap(err)
		if wrapped != nil {
			r.Message = wrapped.Error()
		} else {
			r.Message = err.Error()
		}

		c.JSON(400, r)
	default:
		r.Message = apperror.ErrInternal.Error()
		c.JSON(500, r)
	}
}

func handleSuccess(c *gin.Context, data interface{}) {
	r := delivery.Response{
		Message: "success",
		Data:    data,
	}

	c.JSON(http.StatusOK, r)
}

func (h *handler) Add(c *gin.Context) {
	ctx, _ := context.WithTimeout(c, 3*time.Second)

	var queue Queue
	if err := c.ShouldBindJSON(&queue); err != nil {
		err = fmt.Errorf("Invalid data queue data: %w", err)

		handleError(c, err)
		return
	}

	id, err := h.queueUsecase.Add(ctx, &queue)

	if err != nil {
		handleError(c, err)
		return
	}

	handleSuccess(c, id)
	return
}
