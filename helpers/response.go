package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tamaApotek/tama-go-server/domains/query"
)

func HandleSuccessResponse(c *gin.Context, data interface{}) {
	response := query.Response{
		Message: "success",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func HandleErrorResponse(c *gin.Context, err error) {
	response := query.Response{
		Message: err.Error(),
	}

	var e *query.ErrorQuery
	if errors.As(err, &e) {
		// debugging purpose
		fmt.Print(errors.Unwrap(err))

		response.Code = e.Code()

		c.JSON(http.StatusInternalServerError, response)
	} else {
		c.JSON(http.StatusInternalServerError, response)
	}
}
