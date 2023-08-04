package service

import (
	"context"
	"github.com/agadilkhan/onelab-final/internal/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, u *entity.User) error
}