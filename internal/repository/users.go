package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/v.kirpichov/admin/internal/enitity/models"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) CreateUser(ctx context.Context, user *models.InputUser) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id`)

	fmt.Println(user.Password)
	row := u.db.QueryRowContext(ctx, query, user.Name, user.Password)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UsersRepository) GetUsers(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, name FROM users`
	
	res, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	users := make([]*models.User, 0)

	for res.Next() {
		var user models.User
		if err := res.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if res.Err() != nil {
		return nil, res.Err()
	}

	return users, nil
}
