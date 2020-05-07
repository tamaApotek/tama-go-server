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
	response := Response{
		Message: err.Error(),
	}

	var e *apperror.AppError
	if errors.As(err, &e) {
		// debugging purpose
		fmt.Print(errors.Unwrap(err))

		response.Code = e.Code()

		c.JSON(http.StatusInternalServerError, response)
	} else {
		c.JSON(http.StatusInternalServerError, response)
	}
}
