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
	EmployersRepository EmployerR
	Reports             ReportR

	Sessions      SessionR
	AdminSessions SessionR
}

func NewRepository(db *sqlx.DB, client *redis.Client, adminClient *redis.Client) *Repository {
	return &Repository{
		Users:               NewUsersRepository(db),
		Admins:              NewAdminRepository(db),
		Services:            NewServicesRepository(db),
		EmployersRepository: NewEmployersRepository(db),
		Reports:             NewReportsRepository(db),

		Sessions:      NewSessionRepository(client),
		AdminSessions: NewSessionRepository(adminClient),
	}
}

type UserR interface {
	CreateUser(ctx context.Context, user *models.InputUser) (int64, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAppliesByUser(ctx context.Context, userId int64) ([]*models.Application, error)
}

type AdminR interface {
	GetAdmin(ctx context.Context, username, password string) (*models.Admin, error)
	GetAllApplies(ctx context.Context) ([]*models.ApplyWithData, error)
}

type SessionR interface {
	CreateSession(value, token string, exirationTime time.Duration) error
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}

type ServiesR interface {
	GetService(ctx context.Context, id int64) (*models.Service, error)
	GetServices(ctx context.Context) ([]*models.Service, error)
	Apply(ctx context.Context, userId, serviceId int64) (int64, error)
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

type ReportR interface {
	GetReport(ctx context.Context, id int64) (*models.Report, error)
	GetReports(ctx context.Context) ([]*models.Report, error)
	CreateReport(ctx context.Context) (int64, error)
	DeleteReport(ctx context.Context, id int64) error
}
