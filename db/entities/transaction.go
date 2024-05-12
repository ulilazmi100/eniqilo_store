package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type (
	TransactionPayload struct {
		CustomerId     *string          `db:"customer_id" json:"customerId" validate:"required" `
		Paid           *int             `db:"paid" json:"paid" validate:"required" `
		Change         *int             `db:"change" json:"change" validate:"required" `
		ProductDetails *[]ProductDetail `db:"prodcut_details" json:"productDetails" validate:"required" `
	}

	Transaction struct {
		Id             string           `db:"id" json:"transactionId"`
		CustomerId     *string          `db:"customer_id" json:"customerId" validate:"required" `
		Paid           *int             `db:"paid" json:"paid" validate:"required" `
		Change         *int             `db:"change" json:"change" validate:"required" `
		ProductDetails *[]ProductDetail `db:"prodcut_details" json:"productDetails" validate:"required" `
		CreatedAt      string           `db:"created_at" json:"createdAt"`
	}

	ProductDetail struct {
		ProductId *string `db:"product_id" json:"productId" validate:"required" `
		Quantity  *int    `db:"quantity" json:"quantity" validate:"required" `
	}

	FilterGetTransactions struct {
		CustomerId string `json:"customerId"`
		Limit      int    `json:"limit"`
		Offset     int    `json:"offset"`
		CreatedAt  string `json:"createdAt"`
	}
)

func (t *TransactionPayload) Validate() error {
	err := validation.ValidateStruct(t,
		validation.Field(&t.Paid,
			validation.Required.Error("paid is required"),
			validation.Min(1),
		),
		validation.Field(&t.Change,
			validation.Min(0),
		),
	)

	return err
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
