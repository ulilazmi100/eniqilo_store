package entities

import (
	"errors"
	"regexp"
	"time"

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
		IsAvailable bool      `db:"is_available" json:"isAvailable"`
		CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	}

	ProductRegPayload struct {
		Name        string `db:"name" json:"name"`
		Sku         string `db:"sku" json:"sku"`
		Category    string `db:"category" json:"category"`
		ImageUrl    string `db:"image_url" json:"imageUrl"`
		Notes       string `db:"notes" json:"notes"`
		Price       int    `db:"price" json:"price"`
		Stock       int    `db:"stock" json:"stock"`
		Location    string `db:"location" json:"location"`
		IsAvailable *bool  `db:"is_available" json:"isAvailable" validate:"required" `
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
		Id          string `db:"id" json:"id"`
		Name        string `db:"name" json:"name"`
		Sku         string `db:"sku" json:"sku"`
		Category    string `db:"category" json:"category"`
		ImageUrl    string `db:"image_url" json:"imageUrl"`
		Notes       string `db:"notes" json:"notes"`
		Price       int    `db:"price" json:"price"`
		Stock       int    `db:"stock" json:"stock"`
		Location    string `db:"location" json:"location"`
		IsAvailable bool   `db:"is_available" json:"isAvailable"`
		CreatedAt   string `db:"createdAt" json:"createdAt"`
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

	PaginationMeta struct {
		Limit  int `json:"limit"  validate:"numeric,min=0" schema:"limit"`
		Offset int `json:"offset"  validate:"numeric,min=0" schema:"offset"`
		Total  int `json:"total"`
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
			validation.By(validateUrlFormat),
		),
		validation.Field(&u.Price,
			validation.Required.Error("price is required"),
			validation.Min(1),
		),
		validation.Field(&u.Stock,
			validation.Required.Error("stock is required"),
			validation.Min(0),
			validation.Max(100000),
		),
		validation.Field(&u.Location,
			validation.Required.Error("location is required"),
			validation.Length(1, 200).Error("location must be between 1 and 200 characters"),
		),
		// validation.Field(&u.IsAvailable,
		// 	validation.Required.Error("isAvailable is required"),
		// 	// validation.In("true", "false"),
		// ),
	)

	return err
}

func validateUrlFormat(value any) error {
	url, ok := value.(string)
	if !ok {
		return errors.New("parse error")
	}

	pattern := `(https?:\/\/(?:www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|[a-zA-Z0-9]+\.[^\s]{2,})|www\.[a-zA-Z0-9]+\.[^\s]{2,})`
	rgx := regexp.MustCompile(pattern)
	if !rgx.MatchString(url) {
		return errors.New("invalid Url format")
	}

	return nil
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
