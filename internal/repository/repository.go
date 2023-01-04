package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users UsersRepository
	Sessions SessionRepository
}

func NewRepository(db *sqlx.DB, client *redis.Client) *Repository {
	return &Repository{
		Users: *NewUsersRepository(db),
		Sessions: *NewSessionRepository(client),
	}
}
