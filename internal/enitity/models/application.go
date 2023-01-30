package models

type Application struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"userId"`
	ServiceId int64 `json:"serviceId"`
}
