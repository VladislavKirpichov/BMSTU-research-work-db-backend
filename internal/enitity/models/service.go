package models

type (
	Service struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`
		Cost        int64  `json:"cost"`
		Description string `json:"description"`
	}
)
