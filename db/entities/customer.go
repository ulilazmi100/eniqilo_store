package models

import "time"

type (
	Customer struct {
		Id        string    `db:"id" json:"id"`
		Phone     string    `db:"phone" json:"phoneNumber"`
		Name      string    `db:"name" json:"name"`
		CreatedAt time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	}

	CustomerFilter struct {
		Phone string `json:"phoneNumber"`
		Name  string `json:"name"`
	}
)
