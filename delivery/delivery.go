package delivery

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/domain/apperror"
)

// Delivery provide http response handler helper
type Delivery struct {
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type MyHandler func(c *gin.Context) (interface{}, error)

func handleSuccessResponse(c *gin.Context, data interface{}) {
	r := Response{
		Message: "success",
		Data:    data,
	}

	c.JSON(http.StatusOK, r)
}

func handleErrorResponse(c *gin.Context, err error) {
	fmt.Printf("[ERROR] %+v", err)

	switch {
	case errors.Is(err, apperror.ErrInvalid):
		c.JSON(400, gin.H{"message": apperror.ErrInvalid.Error()})
	default:
		c.JSON(500, gin.H{"message": apperror.ErrInternal.Error()})
	}
}

// Handle will handle http query response
func (d *Delivery) Handle(fn MyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := fn(c)

		if err != nil {
			handleErrorResponse(c, err)
			return
		}

		handleSuccessResponse(c, data)
	}
}
