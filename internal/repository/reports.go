package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

type ReportsRepository struct {
	db *sqlx.DB
}

func NewReportsRepository(db *sqlx.DB) *ReportsRepository {
	return &ReportsRepository{
		db: db,
	}
}

func (r *ReportsRepository) GetReport(ctx context.Context, id int64) (*models.Report, error) {
	query := `SELECT id, created_date, updated_date, leads FROM reports WHERE reports.id=$1`

	row := r.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	report := &models.Report{}
	if err := row.Scan(&report.Id, &report.CreatedDate, &report.UpdatedDate, &report.Leads); err != nil {
		return nil, err
	}

	return report, nil
}

func (r *ReportsRepository) GetReports(ctx context.Context) ([]*models.Report, error) {
	query := `SELECT id, created_date, updated_date, leads FROM reports`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reports := make([]*models.Report, 0)

	for rows.Next() {
		var report models.Report
		if err := rows.Scan(&report.Id, &report.CreatedDate, &report.UpdatedDate, &report.Leads); err != nil {
			return nil, err
		}

		reports = append(reports, &report)
	}

	return reports, nil
}

func (r *ReportsRepository) CreateReport(ctx context.Context) (int64, error) {
	query := `INSERT INTO reports (leads) VALUES ((SELECT COUNT(id) FROM applies)) RETURNING id`

	row := r.db.QueryRowContext(ctx, query)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ReportsRepository) DeleteReport(ctx context.Context, id int64) error {
	query := `DELETE FROM reports WHERE reports.id=$1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("No rows afffected")
	}

	return nil
}
