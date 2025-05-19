package application

import (
	"github.com/SimonKimDev/CoffeeChat/internal/domain/blog"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/db"
)

type BlogPoster interface {
	Post() blog.Post
}

type postService struct {
}

func NewPostService() BlogPoster {
	return &postService{}
}

func (*postService) Post() blog.Post {
	query := `
        INSERT INTO blog.posts (AuthorId, CategoryId, Title, Slug, Summary, BodyMarkdown, DatePublished, DateUpdated)
        VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)            
    `
	db.DB.Prepare(query)
	return blog.Post{}
}
