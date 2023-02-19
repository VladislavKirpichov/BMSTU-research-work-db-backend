package models

type (
	Apply struct {
		Id        int64 `json:"id"`
		UserId    int64 `json:"user_id"`
		ServiceId int64 `json:"service_id"`
	}

	ApplyWithData struct {
		Id          int64  `json:"id"`
		UserId      int64  `json:"user_id"`
		Email       string `json:"email"`
		ServiceId   int64  `json:"service_id"`
		ServiceName string `json:"service_name"`
	}
)
