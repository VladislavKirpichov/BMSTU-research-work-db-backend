package usecase

import (
	"context"

	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
)

type Usecases struct {
	UserUsecase    UserU
	AdminUsecase   AdminU
	ServiceUsecase ServicesU
}

func NewUsecases(repository *repository.Repository, cfg *configs.Config) *Usecases {
	return &Usecases{
		UserUsecase:    NewUserUsecase(repository.Users, repository.Sessions, cfg),
		AdminUsecase:   NewAdminUsecase(repository.Admins, repository.AdminSessions, cfg),
		ServiceUsecase: NewServiceUsecase(repository.Services),
	}
}

type SessionUsecase interface {
	GetSessionToken(ctx context.Context, key string) (string, error)
	Logout(ctx context.Context, key string) error
}

type UserU interface {
	SessionUsecase

	SignIn(ctx context.Context, user *models.SignInUser) (*models.User, error)
	SignUp(ctx context.Context, user *models.InputUser) (int64, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type AdminU interface {
	SessionUsecase

	SignIn(ctx context.Context, admin *models.Admin) error
}

type ServicesU interface {
	GetService(ctx context.Context, id int64) (*models.Service, error)
	CreateService(ctx context.Context, service *models.Service) (int64, error)
	UpdateService(ctx context.Context, service *models.Service) error
	DeleteService(ctx context.Context, id int64) error
}
