package models

type InputUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
