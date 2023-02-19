package models

type (
	SignInUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	InputUser struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	User struct {
		Id       int64  `json:"id" db:"id"`
		Email    string `json:"email" db:"email"`
		Name     string `json:"name" db:"name"`
		Password string `json:"password" db:"password"`
	}
)
