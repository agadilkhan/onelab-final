package pgrepo

import (
	"context"
	"fmt"

	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/agadilkhan/onelab-final/internal/repository"
)

type UserRepository struct {
	Client *Postgres
}

func NewPostgresUser(client *Postgres) repository.UserRepository {
	return &UserRepository{
		Client: client,
	}
}

const usersTable = "users"

func (r *UserRepository) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
					INSERT INTO %s (
						username, --1
						first_name, --2
						last_name, --3
						hashed_password --4
						)
					VALUES ($1, $2, $3, $4)
					`, usersTable)
	_, err := r.Client.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}
