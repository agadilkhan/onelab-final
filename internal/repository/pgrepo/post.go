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
							VALUES ($1, $2, $3, $4)`, postsTable)

	_, err := r.Client.Pool.Exec(ctx, query, p.Title, p.Content, p.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) GetPostByID(ctx context.Context, id int64) (*entity.Post, error) {
	post := new(entity.Post)

	query := fmt.Sprintf(`SELECT title, content, user_id from %s WHERE id=$1`, postsTable)

	err := pgxscan.Get(ctx, r.Client.Pool, post, query, id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) DeletePost(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 RETURNING id, title, content, user_id`, postsTable)

	var deletedPost entity.Post

	err := r.Client.Pool.QueryRow(ctx, query, id).Scan(
		&deletedPost.ID,
		deletedPost.Title,
		deletedPost.Content,
		deletedPost.UserID,
	)
	if err != nil {
		return err
	}

	return nil
}
