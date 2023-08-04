package service

import "github.com/agadilkhan/onelab-final/internal/repository"

type Service struct {
	UserService
	PostService
}

func New(repo repository.Repository) *Service {
	return &Service{
		UserService: repo.UserRepository,
		PostService: repo.PostRepository,
	}
}