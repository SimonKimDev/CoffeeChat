package http

import (
	"encoding/json"
	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"net/http"
)

type PostHandler struct {
	poster application.BlogPoster
}

func NewPostHandler(p application.BlogPoster) *PostHandler {
	return &PostHandler{poster: p}
}

func (p *PostHandler) post(w http.ResponseWriter, r *http.Request) {
	blogPost := p.poster.Post()
	_ = json.NewEncoder(w).Encode(blogPost)
}
