package usecase

import (
	"context"

	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
)

type EmployersUsecases struct {
	repo repository.EmployerR
}

func NewEmployersUsecases(repo repository.EmployerR) *EmployersUsecases {
	return &EmployersUsecases{
		repo: repo,
	}
}

func (e *EmployersUsecases) GetEmployer(ctx context.Context, id int64) (*models.Employer, error) {
	employer, err := e.repo.GetEmployer(ctx, id)
	if err != nil {
		return nil, err
	}

	return employer, nil
}

func (e *EmployersUsecases) GetEmployers(ctx context.Context) ([]*models.Employer, error) {
	employers, err := e.repo.GetEmployers(ctx)
	if err != nil {
		return nil, err
	}

	return employers, nil
}

func (e *EmployersUsecases) CreateEmployer(ctx context.Context, employer *models.Employer) (int64, error) {
	id, err := e.repo.CreateEmployer(ctx, employer)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (e *EmployersUsecases) UpdateEmployer(ctx context.Context, employer *models.Employer) error {
	err := e.repo.UpdateEmployer(ctx, employer)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployersUsecases) DeleteEmployer(ctx context.Context, id int64) error {
	err := e.repo.DeleteEmployer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
