package application

import (
	"context"

	"github.com/SimonKimDev/CoffeeChat/internal/domain/blog"
	"github.com/SimonKimDev/CoffeeChat/internal/infra/db"
)

type BlogPoster interface {
	CreatePost(context.Context, *blog.Post) error
	GetAllPost(context.Context) (*[]blog.Post, error)
	GetPostById(context.Context, int64) (*blog.Post, error)
}

type postService struct {
}

func NewPostService() BlogPoster {
	return &postService{}
}

func (*postService) CreatePost(ctx context.Context, post *blog.Post) error {
	const query = `
        INSERT INTO blog.post(author_id, category_id, title, slug, summary, body_markdown, date_published, date_updated)
        VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
        RETURNING post_id;`

	// common practice is to use Exec, but postgres doesn't allow LastInsertedId
	err := db.DB.QueryRowContext(ctx, query,
		post.AuthorId,
		post.CategoryId,
		post.Title,
		post.Slug,
		post.Summary,
		post.BodyMarkdown,
		post.DatePublished,
	).Scan(&post.PostId)

	return err
}

func (*postService) GetAllPost(ctx context.Context) (*[]blog.Post, error) {
	// TODO: this needs to be paginated
	const query = `select * from blog.post`

	rows, err := db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []blog.Post

	for rows.Next() {
		var post blog.Post
		err := rows.Scan(&post.PostId, &post.AuthorId, &post.CategoryId, &post.Title, &post.Slug, &post.Summary, &post.BodyMarkdown, &post.DatePublished, &post.DateUpdated)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

func (*postService) GetPostById(ctx context.Context, postId int64) (*blog.Post, error) {
	const query = `select * from blog.post where post_id = $1`

	row := db.DB.QueryRowContext(ctx, query, postId)

	var post blog.Post
	err := row.Scan(&post.PostId, &post.AuthorId, &post.CategoryId, &post.Title, &post.Slug, &post.Summary, &post.BodyMarkdown, &post.DatePublished, &post.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
