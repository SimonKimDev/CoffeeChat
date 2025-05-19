package blog

import (
	"database/sql"
	"time"
)

type Post struct {
	PostId        int64          `json:"post_id"`
	AuthorId      int64          `json:"author_id"`
	CategoryId    sql.NullInt64  `json:"category_id"`
	Title         string         `json:"title"`
	Slug          string         `json:"slug"`
	Summary       sql.NullString `json:"summary"`
	BodyMarkdown  string         `json:"body_markdown"`
	DatePublished time.Time      `json:"date_published"`
	DateUpdated   time.Time      `json:"date_updated"`
}
