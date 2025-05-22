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

func RegisterPostRoutes(mux *http.ServeMux, postSvc application.BlogPoster) http.Handler {
	handler := NewPostHandler(postSvc)

	mux.HandleFunc("POST /blogpost", handler.createPost)
	mux.HandleFunc("GET /blogpost", handler.getPosts)
	mux.HandleFunc("GET /blogpost/{id}", handler.getPostById)
	return mux
}
