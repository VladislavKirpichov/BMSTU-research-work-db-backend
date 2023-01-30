package models

import "time"

type Report struct {
	Id          int64     `json:"id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Leads       uint      `json:"leads"`
}
