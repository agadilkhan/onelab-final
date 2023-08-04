package service

import (
	"github.com/agadilkhan/onelab-final/internal/config"
	"github.com/agadilkhan/onelab-final/internal/repository"
	"github.com/agadilkhan/onelab-final/pkg/jwttoken"
)

type Service struct {
	UserService
	PostService
}

func New(repo repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Service {
	return &Service{
		UserService: NewUserService(repo.UserRepository, config, token),
		PostService: NewPostService(repo.PostRepository),
	}
}
