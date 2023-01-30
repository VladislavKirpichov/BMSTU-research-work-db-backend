package usecase

import (
	"context"

	"github.com/v.kirpichov/admin/internal/enitity/models"
	"github.com/v.kirpichov/admin/internal/repository"
)

type ReportsUsecase struct {
	repo repository.ReportR
}

func NewReportsUsecase(repo repository.ReportR) *ReportsUsecase {
	return &ReportsUsecase{
		repo: repo,
	}
}

func (u *ReportsUsecase) Get(ctx context.Context, id int64) (*models.Report, error) {
	return u.repo.GetReport(ctx, id)
}

func (u *ReportsUsecase) GetReports(ctx context.Context) ([]*models.Report, error) {
	reports, err := u.repo.GetReports(ctx)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (u *ReportsUsecase) Create(ctx context.Context) (int64, error) {
	id, err := u.repo.CreateReport(ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}
