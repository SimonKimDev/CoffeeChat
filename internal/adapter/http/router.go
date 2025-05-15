package http

import (
	"net/http"

	"github.com/SimonKimDev/CoffeeChat/internal/application"
)

func RegisterRoutes(mux *http.ServeMux, greetSvc application.Greeter) http.Handler {
	handler := NewGreetingHandler(greetSvc)

	mux.HandleFunc("GET /hello", handler.greet)
	return mux
}

func RegisterUserRoutes(mux *http.ServeMux, userSvc application.BlogPoster) http.Handler {
	handler := New
}
