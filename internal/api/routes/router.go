package routes

import (
	"github.com/SimonKimDev/CoffeeChat/internal/api/handlers"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"net/http"
)

func RegisterPostRoutes(mux *http.ServeMux, postSvc application.BlogPoster) http.Handler {
	handler := handlers.NewPostHandler(postSvc)

	mux.HandleFunc("POST /blogpost", handler.CreatePost)
	mux.HandleFunc("GET /blogpost", handler.GetPosts)
	mux.HandleFunc("GET /blogpost/{id}", handler.GetPostByID)
	mux.HandleFunc("GET /blogpost/category/{id}", handler.GetPostByCategoryID)
	mux.HandleFunc("POST /blogpost/update", handler.UpdatePost)
	mux.HandleFunc("POST /blogpost/delete/{id}", handler.DeletePostByID)
	return mux
}
