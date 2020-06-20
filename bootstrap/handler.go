package bootstrap

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/tamaApotek/tama-go-server/common"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type Handler struct{}

func (h *Handler) HandleErrorResponse(w http.ResponseWriter, err error) {
	r := Response{
		Message: "failed",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	switch {
	case
		errors.Is(err, common.ErrInvalid),
		errors.Is(err, common.ErrInvalidBody):

		w.WriteHeader(400)

		wrapped := errors.Unwrap(err)
		if wrapped != nil {
			r.Error = wrapped.Error()
		} else {
			r.Error = err.Error()
		}

		json.NewEncoder(w).Encode(r)
	default:
		log.Printf("[ERROR] %+v\n", err)

		w.WriteHeader(500)

		json.NewEncoder(w).Encode(r)
	}
}

func (h *Handler) HandleSuccessResponse(w http.ResponseWriter, data interface{}) {
	r := Response{
		Message: "success",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	json.NewEncoder(w).Encode(r)
}

func HandleErrorResponse(w http.ResponseWriter, err error) {
	r := Response{
		Message: "failed",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	switch {
	case
		errors.Is(err, common.ErrInvalid),
		errors.Is(err, common.ErrInvalidBody):

		w.WriteHeader(400)

		wrapped := errors.Unwrap(err)
		if wrapped != nil {
			r.Error = wrapped.Error()
		} else {
			r.Error = err.Error()
		}

		json.NewEncoder(w).Encode(r)
	default:
		log.Printf("[ERROR] %+v\n", err)

		w.WriteHeader(500)

		json.NewEncoder(w).Encode(r)
	}
}
