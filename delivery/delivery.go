package delivery

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaApotek/tama-go-server/common"
)

// Delivery provide http response handler helper
type Delivery struct {
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type MyHandler func(c *gin.Context) (interface{}, error)

// HandleSuccessResponse handle success response 200 with data
func HandleSuccessResponse(w http.ResponseWriter, data interface{}) {
	r := Response{
		Message: "success",
		Data:    data,
	}

	json.NewEncoder(w).Encode(r)
}

func handleErrorResponse(w http.ResponseWriter, err error) {
	r := Response{
		Message: "failed",
	}

	switch {
	case errors.Is(err, common.ErrInvalid):
		wrapped := errors.Unwrap(err)
		if wrapped != nil {
			r.Error = wrapped.Error()

		} else {
			r.Error = err.Error()
		}

		json.NewEncoder(w).Encode(r)
	default:
		log.Printf("[ERROR] %+v\n", err)

		json.NewEncoder(w).Encode(r)
	}
}

// Handle will handle http query response
func (d *Delivery) Handle(fn MyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := fn(c)

		if err != nil {
			handleErrorResponse(c.Writer, err)
			return
		}

		HandleSuccessResponse(c.Writer, data)
	}
}
