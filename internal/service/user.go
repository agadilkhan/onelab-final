package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/agadilkhan/onelab-final/internal/config"
	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/agadilkhan/onelab-final/internal/repository"
	"github.com/agadilkhan/onelab-final/pkg/jwttoken"
	"github.com/agadilkhan/onelab-final/pkg/util"
)

type UserService interface {
	Register(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, password string) (string, error)
}

type userservice struct {
	Repo repository.UserRepository
	Config *config.Config
	Token *jwttoken.JWTToken
}

func NewUserService(repo repository.UserRepository, config *config.Config, token *jwttoken.JWTToken) *userservice {
	return &userservice{
		Repo: repo,
		Config: config,
		Token: token,
	}
}

func (s *userservice) Register(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = s.Repo.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (s *userservice) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.Repo.GetUser(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := s.Token.GenerateToken(user.ID, s.Config.Token.TTL)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}

func (s *userservice) VerifyToken(token string) (int64, error) {
	payload, err := s.Token.ValidateToken(token)
	if err != nil {
		return 0, fmt.Errorf("validate token err: %w", err)
	}

	return payload.UserID, nil
}