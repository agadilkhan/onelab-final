package pgrepo

import (
	"context"
	"fmt"

	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/georgysavva/scany/pgxscan"
)

type PostRepository struct {
	Client *Postgres
}

func NewPostgresPost(client *Postgres) *PostRepository {
	return &PostRepository{
		Client: client,
	}
}

const postsTable = "posts"

func (r *PostRepository) CreatePost(ctx context.Context, p *entity.Post) error {
	query := fmt.Sprintf(`INSERT INTO %s 
								(title, --1
								content, --2
								user_id --3
								)
							VALUES ($1, $2, $3)`, postsTable)

	_, err := r.Client.Pool.Exec(ctx, query, p.Title, p.Content, p.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) GetPostByID(ctx context.Context, id int) (*entity.Post, error) {
	post := new(entity.Post)

	query := fmt.Sprintf(`SELECT id, title, content, user_id from %s WHERE id=$1`, postsTable)

	err := pgxscan.Get(ctx, r.Client.Pool, post, query, id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) DeletePost(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 RETURNING id, title, content, user_id`, postsTable)

	var deletedPost entity.Post

	r.Client.Pool.QueryRow(ctx, query, id).Scan(
		deletedPost.ID,
		deletedPost.Title,
		deletedPost.Content,
		deletedPost.UserID,
	)

	return nil
}

func (r *PostRepository) GetAllPosts(ctx context.Context) (*[]entity.Post, error) {
	var posts []entity.Post
    query := "SELECT * FROM posts"

    conn, err := r.Client.Pool.Acquire(context.Background())
    if err != nil {
        return nil, err
    }
    defer conn.Release()

    if err := pgxscan.Select(context.Background(), conn, &posts, query); err != nil {
        return nil, err
    }

    return &posts, nil
}

func (r *PostRepository) UpdatePost(ctx context.Context, updatedPost *entity.Post) (*entity.Post, error) {
	query := `
        UPDATE posts
        SET title = $2, content = $3
        WHERE id = $1
    `

    conn, err := r.Client.Pool.Acquire(context.Background())
    if err != nil {
        return nil, err
    }
    defer conn.Release()

    _, err = conn.Exec(context.Background(), query, updatedPost.ID, updatedPost.Title, updatedPost.Content)
    if err != nil {
        return nil, err
    }

    return updatedPost, nil
}