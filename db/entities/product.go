package entities

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	Product struct {
		Id          string    `db:"id" json:"id"`
		Name        string    `db:"name" json:"name"`
		Sku         string    `db:"sku" json:"sku"`
		Category    string    `db:"category" json:"category"`
		ImageUrl    string    `db:"image_url" json:"imageUrl"`
		Notes       string    `db:"notes" json:"notes"`
		Price       int       `db:"price" json:"price"`
		Stock       int       `db:"stock" json:"stock"`
		Location    string    `db:"location" json:"location"`
		IsAvailable bool      `db:"is_avail" json:"isAvailable"`
		CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	}

	ProductRegPayload struct {
		Name        string    `db:"name" json:"name"`
		Sku         string    `db:"sku" json:"sku"`
		Category    string    `db:"category" json:"category"`
		ImageUrl    string    `db:"image_url" json:"imageUrl"`
		Notes       string    `db:"notes" json:"notes"`
		Price       int       `db:"price" json:"price"`
		Stock       int       `db:"stock" json:"stock"`
		Location    string    `db:"location" json:"location"`
		IsAvailable time.Time `db:"is_avail" json:"isAvailable"`
	}

	FilterGetProducts struct {
		Id          string `json:"id"`
		Limit       int    `json:"limit"`
		Offset      int    `json:"offset"`
		Name        string `json:"Name"`
		IsAvailable string `json:"isAvailable"`
		Category    string `json:"category"`
		Sku         string `json:"sku"`
		Price       string `json:"price"`
		InStock     string `json:"inStock"`
		CreatedAt   string `json:"createdAt"`
	}

	ProductList struct {
		Id          string    `db:"id" json:"id"`
		Name        string    `db:"name" json:"name"`
		Sku         string    `db:"sku" json:"sku"`
		Category    string    `db:"category" json:"category"`
		ImageUrl    string    `db:"image_url" json:"imageUrl"`
		Notes       string    `db:"notes" json:"notes"`
		Price       int       `db:"price" json:"price"`
		Stock       int       `db:"stock" json:"stock"`
		Location    string    `db:"location" json:"location"`
		IsAvailable time.Time `db:"is_avail" json:"isAvailable"`
		CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
	}

	FilterSku struct {
		Limit    int    `json:"limit"`
		Offset   int    `json:"offset"`
		Name     string `json:"Name"`
		Category string `json:"category"`
		Sku      string `json:"sku"`
		Price    string `json:"price"`
		InStock  string `json:"inStock"`
	}

	CustomerProductList struct {
		Id        string `db:"id" json:"id"`
		Name      string `db:"name" json:"name"`
		Sku       string `db:"sku" json:"sku"`
		Category  string `db:"category" json:"category"`
		ImageUrl  string `db:"image_url" json:"imageUrl"`
		Price     int    `db:"price" json:"price"`
		Stock     int    `db:"stock" json:"stock"`
		Location  string `db:"location" json:"location"`
		CreatedAt string `db:"createdAt" json:"createdAt"`
	}
)

func (u *ProductRegPayload) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name,
			validation.Required.Error("name is required"),
			validation.Length(1, 30).Error("name must be between 1 and 30 characters"),
		),
		validation.Field(&u.Sku,
			validation.Required.Error("sku is required"),
			validation.Length(1, 30).Error("sku must be between 1 and 30 characters"),
		),
		validation.Field(&u.Category,
			validation.Required.Error("category is required"),
			validation.In("Clothing", "Accessories", "Footwear", "Beverages"),
		),
		validation.Field(&u.Notes,
			validation.Required.Error("notes is required"),
			validation.Length(1, 200).Error("notes must be between 1 and 200 characters"),
		),
		validation.Field(&u.ImageUrl,
			validation.Required.Error("imageUrl is required"),
			is.URL,
		),
		validation.Field(&u.Price,
			validation.Required.Error("price is required"),
			validation.Min(1),
		),
		validation.Field(&u.Stock,
			validation.Required.Error("stock is required"),
			validation.Min(1),
			validation.Max(100000),
		),
		validation.Field(&u.Location,
			validation.Required.Error("location is required"),
			validation.Length(1, 200).Error("location must be between 1 and 200 characters"),
		),
	)

	return err
}

func CategoryChecker(category string) string {
	err := validation.Field(category,
		validation.In("Clothing", "Accessories", "Footwear", "Beverages"),
	)
	if err != nil {
		return ""
	}
	return category
}
