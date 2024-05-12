package entities

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	Customer struct {
		Id          string    `db:"id" json:"id"`
		PhoneNumber string    `db:"phone" json:"phoneNumber"`
		Name        string    `db:"name" json:"name"`
		CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	}

	CustomerRegPayload struct {
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name,omitempty"`
	}

	CustomerFilter struct {
		PhoneNumber string `db:"phone" json:"phoneNumber"`
		Name        string `db:"name" json:"name"`
	}

	CustomerList struct {
		Id          string `db:"id" json:"userId"`
		PhoneNumber string `db:"phone" json:"phoneNumber"`
		Name        string `db:"name" json:"name"`
	}
)

func NewCustomer(phone, name string) *Customer {
	cust := &Customer{
		PhoneNumber: phone,
		Name:        name,
	}

	return cust
}

func (u *Customer) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.PhoneNumber,
			validation.Required.Error("phone is required"),
			validation.Length(10, 16).Error("phone number must be between 10 and 16 characters"),
			validation.By(validatePhoneFormat),
		),
		validation.Field(&u.Name,
			validation.Required.Error("name is required"),
			validation.Length(5, 50).Error("name must be between 5 and 50 characters"),
		),
	)

	return err
}
