package queue

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/delivery"
)

const validationFailed = "Queue validation failed"

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

		r.Message = err.Error()

		wrapped := errors.Unwrap(err)
		if wrapped != nil {
			r.Error = wrapped.Error()

		} else {
			r.Error = err.Error()
		}

		c.JSON(http.StatusBadRequest, r)
	default:
		delivery.HandleErrorResponse(c.Writer, err)
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

		c.JSON(http.StatusBadRequest, delivery.Response{Message: validationFailed, Error: err.Error()})
		return
	}

	id, err := h.queueUsecase.Add(ctx, &queue)

	if err != nil {
		handleError(c, err)
		return
	}

	delivery.HandleSuccessResponse(c.Writer, id)
	return
}
