package usecase

import (
	"context"
	"fmt"

	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
)

type ServiceUsecase struct {
	repo repository.ServiesR
}

func NewServiceUsecase(repo repository.ServiesR) *ServiceUsecase {
	return &ServiceUsecase{
		repo: repo,
	}
}

func (s *ServiceUsecase) GetService(ctx context.Context, id int64) (*models.Service, error) {
	service, err := s.repo.GetService(ctx, id)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func validateService(service *models.Service) error {
	if len([]rune(service.Name)) > 30 {
		return fmt.Errorf("Bad length")
	}

	return nil
}

func (s *ServiceUsecase) CreateService(ctx context.Context, service *models.Service) (int64, error) {
	err := validateService(service)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.CreateService(ctx, service)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (s *ServiceUsecase) UpdateService(ctx context.Context, service *models.Service) error {
	err := validateService(service)
	if err != nil {
		return err
	}

	err = s.repo.UpdateService(ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceUsecase) DeleteService(ctx context.Context, id int64) error {
	err := s.repo.DeleteService(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
