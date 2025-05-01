package api

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(server *http.ServeMux) {
	server.HandleFunc("GET /hello", helloHandler))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
