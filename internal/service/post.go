package service

import (
	"context"

	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/agadilkhan/onelab-final/internal/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, p *entity.Post) error
	GetPostByID(ctx context.Context, id int64) (*entity.Post, error)
	DeletePost(ctx context.Context, id int64) error
}

type postservice struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *postservice {
	return &postservice{
		Repo: repo,
	}
}

func (s *postservice) CreatePost(ctx context.Context, p *entity.Post) error {
	return nil
}

func (s *postservice) GetPostByID(ctx context.Context, id int64) (*entity.Post, error) {
	return nil, nil
}

func (s *postservice) DeletePost(ctx context.Context, id int64) error {
	return nil
}