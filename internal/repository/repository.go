package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

type Repository struct {
	Users         UserR
	Admins        AdminR
	Sessions      SessionR
	AdminSessions SessionR
}

func NewRepository(db *sqlx.DB, client *redis.Client, adminClient *redis.Client) *Repository {
	return &Repository{
		Users:         NewUsersRepository(db),
		Admins:        NewAdminRepository(db),
		Sessions:      NewSessionRepository(client),
		AdminSessions: NewSessionRepository(adminClient),
	}
}

type UserR interface {
	CreateUser(ctx context.Context, user *models.InputUser) (int64, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type AdminR interface {
	GetAdmin(ctx context.Context, username, password string) (*models.Admin, error)
}

type SessionR interface {
	CreateSession(value, token string, exirationTime time.Duration) error
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}
