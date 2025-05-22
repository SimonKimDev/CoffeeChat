package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/SimonKimDev/CoffeeChat/internal/application"
	"github.com/SimonKimDev/CoffeeChat/internal/domain/blog"
)

type PostHandler struct {
	poster application.BlogPoster
}

func NewPostHandler(p application.BlogPoster) *PostHandler {
	return &PostHandler{poster: p}
}

func (p *PostHandler) createPost(w http.ResponseWriter, r *http.Request) {
	const maxBody = 1 << 20 // about 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, maxBody)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var post blog.Post

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&post)
	if err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = p.poster.CreatePost(ctx, &post)

	if err != nil {
		http.Error(w, "failure to create post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if decoder.More() {
		http.Error(w, "request body must contain single JSON Object", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Post is Created")
}

func (p *PostHandler) getPosts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	posts, err := p.poster.GetAllPost(ctx)
	if err != nil {
		http.Error(w, "failure to retrieve posts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		_ = r.Body.Close()
		return
	}
}

func (p *PostHandler) getPostById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "failure to parse id from request url:"+err.Error(), http.StatusBadRequest)
		return
	}

	post, err := p.poster.GetPostById(ctx, id)
	if err != nil {
		http.Error(w, "failure to retrieve post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		_ = r.Body.Close()
		return
	}
}
