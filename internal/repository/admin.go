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
		return err
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
