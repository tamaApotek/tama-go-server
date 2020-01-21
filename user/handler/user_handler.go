package handler

import (
	"fmt"
	"net/http"
)

// UserHandler represent http handler for user
type UserHandler struct {
}

// NewUserHandler will initiate user/ resources endpoint
func NewUserHandler(s *http.ServeMux) {
	handler := &UserHandler{}

	s.HandleFunc("/users", handler.FetchUsers)
}

func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
