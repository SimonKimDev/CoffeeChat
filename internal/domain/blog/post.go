package blog

import "time"

type Post struct {
	PostId        int64 `binding:"required"`
	AuthorId      int64 `binding:"required"`
	CategoryId    int
	Title         string `binding:"required"`
	Slug          string
	Summary       string
	BodyMarkdown  string    `binding:"required"`
	DatePublished time.Time `binding:"required"`
	DateUpdated   time.Time
}
