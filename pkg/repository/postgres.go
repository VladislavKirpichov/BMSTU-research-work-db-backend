package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/v.kirpichov/admin/configs"
)

const driverName = "postgres"
3
func NewPostgresRepository(cfg *configs.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(driverName, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
