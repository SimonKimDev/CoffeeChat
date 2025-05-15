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

func RegisterBlogPostRoutes(mux *http.ServeMux, postSvc application.BlogPoster) http.Handler {
	handler := NewPostHandler(postSvc)

	mux.HandleFunc("GET /user", handler.post)
	return mux
}
