package application

import (
	"github.com/SimonKimDev/CoffeeChat/internal/domain/blog"
)

type BlogPoster interface {
	Post() blog.Post
}

type postService struct {
}

func (*postService) Post() blog.Post {
	return blog.Post{}
}
