package models

type (
	Admin struct {
		Username string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}
)
