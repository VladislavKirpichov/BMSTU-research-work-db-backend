package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

type EmployersRepository struct {
	db *sqlx.DB
}

func NewEmployersRepository(db *sqlx.DB) *EmployersRepository {
	return &EmployersRepository{
		db: db,
	}
}

func (e *EmployersRepository) GetEmployer(ctx context.Context, id int64) (*models.Employer, error) {
	query := `SELECT id, name FROM employers WHERE employers.id=$1`

	row := e.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	employer := &models.Employer{}
	if err := row.Scan(&employer.Id, &employer.Name); err != nil {
		return nil, err
	}

	return employer, nil
}

func (e *EmployersRepository) GetEmployers(ctx context.Context) ([]*models.Employer, error) {
	query := `SELECT id, name FROM employers`

	rows, err := e.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employers := make([]*models.Employer, 0)

	for rows.Next() {
		var employer models.Employer
		if err := rows.Scan(&employer.Id, &employer.Name); err != nil {
			return nil, err
		}

		employers = append(employers, &employer)
	}

	return employers, nil
}

func (e *EmployersRepository) CreateEmployer(ctx context.Context, employer *models.Employer) (int64, error) {
	query := `INSERT INTO employers (name) VALUES ($1) RETURNING id`

	row := e.db.QueryRowContext(ctx, query, employer.Name)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (e *EmployersRepository) UpdateEmployer(ctx context.Context, employer *models.Employer) error {
	query := `UPDATE employers SET name=$1 WHERE employers.id=$2`

	res, err := e.db.ExecContext(ctx, query, employer.Name, employer.Id)
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

func (e *EmployersRepository) DeleteEmployer(ctx context.Context, id int64) error {
	query := `DELETE FROM employers WHERE employers.id=$1`

	res, err := e.db.ExecContext(ctx, query, id)
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
