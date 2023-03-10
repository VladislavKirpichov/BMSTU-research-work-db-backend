package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

type ServicesRepository struct {
	db *sqlx.DB
}

func NewServicesRepository(db *sqlx.DB) *ServicesRepository {
	return &ServicesRepository{
		db: db,
	}
}

func (s *ServicesRepository) GetService(ctx context.Context, id int64) (*models.Service, error) {
	query := `SELECT id, name, cost, description FROM services WHERE services.id=$1`

	row := s.db.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	service := &models.Service{}
	if err := row.Scan(&service.Id, &service.Name, &service.Cost, &service.Description); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *ServicesRepository) GetServices(ctx context.Context) ([]*models.Service, error) {
	query := `SELECT id, name, cost, description FROM services`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	services := make([]*models.Service, 0)

	for rows.Next() {
		var service models.Service
		if err := rows.Scan(&service.Id, &service.Name, &service.Cost, &service.Description); err != nil {
			return nil, err
		}

		services = append(services, &service)
	}

	return services, nil
}

func (s *ServicesRepository) CreateService(ctx context.Context, service *models.Service) (int64, error) {
	query := `INSERT INTO services (name, cost, description) VALUES ($1, $2, $3) RETURNING id`

	row := s.db.QueryRowContext(ctx, query, service.Name, service.Cost, service.Description)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServicesRepository) Apply(ctx context.Context, userId, serviceId int64) (int64, error) {
	query := `INSERT INTO applies (user_id, service_id) VALUES ($1, $2) RETURNING id`

	row := s.db.QueryRowContext(ctx, query, userId, serviceId)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServicesRepository) UpdateService(ctx context.Context, service *models.Service) error {
	query := `UPDATE services
		SET name=$1
		SET cost=$2
		SET description=$3
		WHERE services.id=$4`

	res, err := s.db.ExecContext(ctx, query, service.Name, service.Cost, service.Description,  service.Id)
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

func (s *ServicesRepository) DeleteService(ctx context.Context, id int64) error {
	query := `DELETE FROM services WHERE services.id=$1`

	res, err := s.db.ExecContext(ctx, query, id)
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
