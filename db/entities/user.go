package entities

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id        string    `db:"id" json:"id"`
	Phone     string    `db:"phone" json:"phoneNumber"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"password" json:"password,omitempty"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}

type RegistrationPayload struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name,omitempty"`
	Password    string `json:"password"`
}

type Credential struct {
	Phone    string `json:"phoneNumber"`
	Password string `json:"password"`
}

type JWTPayload struct {
	Id    string
	Phone string
	Name  string
}

type JWTClaims struct {
	Id    string
	Phone string
	Name  string
	jwt.RegisteredClaims
}

func NewUser(phone, name, password string) *User {
	u := &User{
		Phone:    phone,
		Name:     name,
		Password: password,
	}

	return u
}

func (u *Credential) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Phone,
			validation.Required.Error("phone is required"),
			validation.By(validatePhoneFormat),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password is required"),
			validation.Length(5, 15).Error("password must be between 5 and 15 characters"),
		),
	)

	return err
}

func (u *User) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Phone,
			validation.Required.Error("phone is required"),
			validation.Length(10, 16).Error("phone number must be between 10 and 16 characters"),
			validation.By(validatePhoneFormat),
		),
		validation.Field(&u.Name,
			validation.Required.Error("name is required"),
			validation.Length(5, 50).Error("name must be between 5 and 50 characters"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password is required"),
			validation.Length(5, 15).Error("password must be between 5 and 15 characters"),
		),
	)

	return err
}

func validatePhoneFormat(value any) error {
	phone, ok := value.(string)
	if !ok {
		return errors.New("parse error")
	}

	pattern := `^\+((?:9[679]|8[035789]|6[789]|5[90]|42|3[578]|2[1-689])|9[0-58]|8[1246]|6[0-6]|5[1-8]|4[013-9]|3[0-469]|2[70]|7|1)(?:\W*\d){0,}\d$`
	rgx := regexp.MustCompile(pattern)
	if !rgx.MatchString(phone) {
		return errors.New("invalid phone format")
	}

	return nil
}
