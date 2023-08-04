package service

import (
	"context"

	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/agadilkhan/onelab-final/internal/repository"
)

type PostService interface {
	GetAllPosts(ctx context.Context) (*[]entity.Post, error)
	CreatePost(ctx context.Context, p *entity.Post) error
	GetPostByID(ctx context.Context, id int) (*entity.Post, error)
	DeletePost(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, p *entity.Post) (*entity.Post, error)
}

type postservice struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *postservice {
	return &postservice{
		Repo: repo,
	}
}

func (s *postservice) GetAllPosts(ctx context.Context) (*[]entity.Post, error) {
	return s.Repo.GetAllPosts(ctx)
}

func (s *postservice) CreatePost(ctx context.Context, p *entity.Post) error {
	return s.Repo.CreatePost(ctx, p)
}

func (s *postservice) GetPostByID(ctx context.Context, id int) (*entity.Post, error) {
	return s.Repo.GetPostByID(ctx, id)
}

func (s *postservice) DeletePost(ctx context.Context, id int) error {
	return s.Repo.DeletePost(ctx, id)
}

func (s *postservice) UpdatePost(ctx context.Context, p *entity.Post) (*entity.Post, error) {
	return s.Repo.UpdatePost(ctx, p)
}