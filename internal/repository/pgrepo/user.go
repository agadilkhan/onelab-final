package pgrepo

import (
	"context"
	"fmt"
	"strings"

	"github.com/agadilkhan/onelab-final/internal/entity"
	"github.com/georgysavva/scany/pgxscan"
)

type UserRepository struct {
	Client *Postgres
}

func NewPostgresUser(client *Postgres) *UserRepository {
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

func (r *UserRepository) GetUser(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)

	query := fmt.Sprintf(`SELECT first_name, last_name, username, hashed_password FROM %s WHERE username=$1`, usersTable)

	err := pgxscan.Get(ctx, r.Client.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}