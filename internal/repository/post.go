package repository

import (
	"context"
	"github.com/agadilkhan/onelab-final/internal/entity"
)

type PostRepository interface {
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	CreatePost(ctx context.Context, p *entity.Post) error
	GetPostByID(ctx context.Context, id int64) (*entity.Post, error)
	DeletePost(ctx context.Context, id int64) error
}