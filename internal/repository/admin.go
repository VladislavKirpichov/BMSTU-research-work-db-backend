package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

// TODO: вынести в конфиг (в идеале в vault)
const (
	adminUsername = "admin"
	adminPassword = "admin"
)

type AdminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

// TODO: продумать инициализацию админов
func MustInitAdmins(db *sqlx.DB) error {
	query := `INSERT INTO admins (username, password) VALUES ($1, $2)`

	err := db.QueryRow(query, adminUsername, adminPassword).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (a *AdminRepository) GetAdmin(ctx context.Context, username, password string) (*models.Admin, error) {
	query := `SELECT username, password FROM admins WHERE admins.username=$1 and admins.password=$2`

	row := a.db.QueryRowxContext(ctx, query, username, password)
	if row.Err() != nil {
		return nil, row.Err()
	}

	admin := &models.Admin{}

	err := row.StructScan(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (a *AdminRepository) GetAllApplies(ctx context.Context) ([]*models.ApplyWithData, error) {
	query := `SELECT applies.id, applies.user_id, users.email, applies.service_id, services.name
		FROM applies
	    JOIN users ON users.id=applies.user_id
		JOIN services ON services.id=applies.service_id`

	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applies := make([]*models.ApplyWithData, 0)

	for rows.Next() {
		var apply models.ApplyWithData
		if err := rows.Scan(&apply.Id, &apply.UserId, &apply.Email, &apply.ServiceId, &apply.ServiceName); err != nil {
			return nil, err
		}

		applies = append(applies, &apply)
	}

	return applies, nil
}
