package main

import (
	"net/http"

	"github.com/tamaApotek/tama-go-server/user/handler"
)

func main() {
	mux := http.NewServeMux()
	handler.NewUserHandler(mux)
}
