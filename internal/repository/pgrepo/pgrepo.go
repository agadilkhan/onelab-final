package pgrepo

import "github.com/agadilkhan/onelab-final/internal/repository"

func NewPostgresRepository(client *Postgres) *repository.Repository {
	return &repository.Repository{
		UserRepository: NewPostgresUser(client),
		PostRepository: NewPostgresPost(client),
	}
}