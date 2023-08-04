package repository

import (
	"context"
	"github.com/agadilkhan/onelab-final/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.User) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
}