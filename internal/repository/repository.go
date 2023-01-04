package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Users UsersRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: *NewUsersRepository(db),
	}
}
