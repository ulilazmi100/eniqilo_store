package models

import "time"

type User struct {
	Id        string    `db:"id" json:"id"`
	Phone     string    `db:"phone" json:"phoneNumber"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"password" json:"password,omitempty"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}
