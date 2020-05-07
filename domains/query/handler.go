package query

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/domains/apperror"
)

func HandleSuccessResponse(c *gin.Context, data interface{}) {
	response := Response{
		Message: "success",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func HandleErrorResponse(c *gin.Context, err error) {
	switch {
	case errors.Is(err, apperror.ErrInvalid):
		c.JSON(400, gin.H{"message": apperror.ErrInvalid.Error()})
	default:
		fmt.Printf("[ERROR] %+v", err)
		c.JSON(500, gin.H{"message": apperror.ErrInternal.Error()})
	}
}
