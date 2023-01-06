package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

type Repository struct {
	Users               UserR
	Admins              AdminR
	Services            ServiesR
	Sessions            SessionR
	AdminSessions       SessionR
	EmployersRepository EmployerR
}

func NewRepository(db *sqlx.DB, client *redis.Client, adminClient *redis.Client) *Repository {
	return &Repository{
		Users:               NewUsersRepository(db),
		Admins:              NewAdminRepository(db),
		Services:            NewServicesRepository(db),
		EmployersRepository: NewEmployersRepository(db),

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

type ServiesR interface {
	GetService(ctx context.Context, id int64) (*models.Service, error)
	GetServices(ctx context.Context) ([]*models.Service, error)
	CreateService(ctx context.Context, service *models.Service) (int64, error)
	UpdateService(ctx context.Context, service *models.Service) error
	DeleteService(ctx context.Context, id int64) error
}

type EmployerR interface {
	GetEmployer(ctx context.Context, id int64) (*models.Employer, error)
	GetEmployers(ctx context.Context) ([]*models.Employer, error)
	CreateEmployer(ctx context.Context, employer *models.Employer) (int64, error)
	UpdateEmployer(ctx context.Context, employer *models.Employer) error
	DeleteEmployer(ctx context.Context, id int64) error
}
