package handlers

import (
	"context"
	"encoding/json"
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

func (p *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	const maxBody = 1 << 20 // about 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, maxBody)

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var post blog.Post

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&post)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = p.poster.CreatePost(ctx, &post)

	if err != nil {
		http.Error(w, "Failed to create post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if decoder.More() {
		http.Error(w, "Request body must contain single JSON Object", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	posts, err := p.poster.GetAllPost(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve posts: "+err.Error(), http.StatusInternalServerError)
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

func (p *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse id from request url:"+err.Error(), http.StatusBadRequest)
		return
	}

	post, err := p.poster.GetPostByID(ctx, id)
	if err != nil {
		http.Error(w, "Failed to retrieve post: "+err.Error(), http.StatusInternalServerError)
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

func (p *PostHandler) GetPostByCategoryID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse id from request url:"+err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := p.poster.GetPostByCategoryID(ctx, id)
	if err != nil {
		http.Error(w, "Failed to retrieve posts by category ID"+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		_ = r.Body.Close()
		return
	}
}

func (p *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var post blog.Post
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&post)
	if err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = p.poster.UpdatePost(ctx, &post)

	if err != nil {
		http.Error(w, "Failed to update post: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PostHandler) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Failed to parse postid:"+err.Error(), http.StatusBadRequest)
	}

	err = p.poster.DeletePostByID(ctx, id)

	if err != nil {
		http.Error(w, "Failed to delete post: "+err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
