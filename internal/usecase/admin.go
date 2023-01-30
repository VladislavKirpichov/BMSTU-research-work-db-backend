package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/v.kirpichov/admin/configs"
	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
)

type AdminUsecase struct {
	cfg          *configs.Config
	adminsRepo   repository.AdminR
	sessionsRepo repository.SessionR
}

func NewAdminUsecase(adminRepo repository.AdminR, sessionRepo repository.SessionR, cfg *configs.Config) *AdminUsecase {
	return &AdminUsecase{
		adminsRepo:   adminRepo,
		sessionsRepo: sessionRepo,
		cfg:          cfg,
	}
}

func (a *AdminUsecase) SignIn(ctx context.Context, admin *models.Admin) error {
	_, err := a.adminsRepo.GetAdmin(ctx, admin.Username, admin.Password)
	if err != nil {
		return err
	}

	return nil
}

func (a *AdminUsecase) GetSessionToken(ctx context.Context, username string) (string, error) {
	token := uuid.NewString()

	err := a.sessionsRepo.CreateSession(username, token, time.Duration(a.cfg.AdminSessionConfig.ExpiresAt*int(time.Hour.Nanoseconds())))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AdminUsecase) Auth(ctx context.Context, token string) (*models.Admin, error) {
	username, err := a.sessionsRepo.GetSession(token)
	if err != nil {
		return nil, err
	}

	return &models.Admin{
		Username: username,
	}, nil
}

func (a *AdminUsecase) Logout(ctx context.Context, username string) error {
	return a.sessionsRepo.DeleteSession(username)
}

func (a *AdminUsecase) GetAllApplies(ctx context.Context) ([]*models.ApplyWithData, error) {
	applies, err := a.adminsRepo.GetAllApplies(ctx)
	if err != nil {
		return nil, err
	}

	return applies, nil
}
