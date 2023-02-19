package usecase

import (
	"context"

	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
)

type Usecases struct {
	UserUsecase       UserU
	AdminUsecase      AdminU
	ServiceUsecase    ServicesU
	EmployersUsecases EmployerU
	ReportsUsecase    ReportU
}

func NewUsecases(repository *repository.Repository, cfg *configs.Config) *Usecases {
	return &Usecases{
		UserUsecase:       NewUserUsecase(repository.Users, repository.Sessions, cfg),
		AdminUsecase:      NewAdminUsecase(repository.Admins, repository.AdminSessions, cfg),
		ServiceUsecase:    NewServiceUsecase(repository.Services),
		EmployersUsecases: NewEmployersUsecases(repository.EmployersRepository),
		ReportsUsecase:    NewReportsUsecase(repository.Reports),
	}
}

type SessionUsecase interface {
	GetSessionToken(ctx context.Context, key string) (string, error)
	Logout(ctx context.Context, key string) error
}

type UserU interface {
	SessionUsecase

	Auth(ctx context.Context, token string) (*models.User, error)
	SignIn(ctx context.Context, user *models.SignInUser) (*models.User, error)
	SignUp(ctx context.Context, user *models.InputUser) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)

	GetAppliesByUser(ctx context.Context, userId int64) ([]*models.Application, error)
}

type AdminU interface {
	SessionUsecase
	Auth(ctx context.Context, token string) (*models.Admin, error)
	SignIn(ctx context.Context, admin *models.Admin) error
	GetAllApplies(ctx context.Context) ([]*models.ApplyWithData, error)
}

type ServicesU interface {
	GetService(ctx context.Context, id int64) (*models.Service, error)
	GetServices(ctx context.Context) ([]*models.Service, error)
	Apply(ctx context.Context, userId, serviceId int64) (int64, error)
	CreateService(ctx context.Context, service *models.Service) (int64, error)
	UpdateService(ctx context.Context, service *models.Service) error
	DeleteService(ctx context.Context, id int64) error
}

type EmployerU interface {
	GetEmployer(ctx context.Context, id int64) (*models.Employer, error)
	GetEmployers(ctx context.Context) ([]*models.Employer, error)
	CreateEmployer(ctx context.Context, service *models.Employer) (int64, error)
	UpdateEmployer(ctx context.Context, service *models.Employer) error
	DeleteEmployer(ctx context.Context, id int64) error
}

type ReportU interface {
	Create(ctx context.Context) (int64, error)
	GetReports(ctx context.Context) ([]*models.Report, error)
	Get(ctx context.Context, id int64) (*models.Report, error)
}
