package entities

import (
	"time"
)

type (
	Product struct {
		Id        string    `db:"id" json:"id"`
		Name      string    `db:"name" json:"name"`
		Sku       string    `db:"sku" json:"sku"`
		ImageUrl  string    `db:"image_url" json:"imageUrl"`
		Notes     string    `db:"notes" json:"notes"`
		Price     int       `db:"price" json:"price"`
		Stock     int       `db:"stock" json:"stock"`
		Location  string    `db:"location" json:"location"`
		IsAvail   time.Time `db:"is_avail" json:"isAvailable"`
		CreatedAt time.Time `db:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
	}

	FilterGetProducts struct {
		Id        string `json:"id"`
		Limit     int    `json:"limit"`
		Offset    int    `json:"offset"`
		Name      string `json:"Name"`
		IsAvail   string `json:"isAvailable"`
		Category  string `json:"category"`
		Sku       string `json:"sku"`
		Price     int    `json:"price"`
		InStock   bool   `json:"inStock"`
		CreatedAt string `json:"createdAt"`
	}

	FilterSku struct {
		Limit    int    `json:"limit"`
		Offset   int    `json:"offset"`
		Name     string `json:"Name"`
		Category string `json:"category"`
		Sku      string `json:"sku"`
		Price    int    `json:"price"`
		InStock  bool   `json:"inStock"`
	}
)
