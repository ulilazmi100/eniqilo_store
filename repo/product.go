package repo

import (
	"eniqilo_store/db/entities"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo interface {
	AddProduct(ctx *fiber.Ctx, product *entities.ProductRegPayload) (string, time.Time, error)
	SearchProduct(ctx *fiber.Ctx, filter entities.FilterGetProducts) ([]entities.ProductList, error)
	SearchProductCustomer(ctx *fiber.Ctx, filter entities.FilterSku) ([]entities.CustomerProductList, error)
	GetProductById(ctx *fiber.Ctx, id string) (*entities.Product, error)
	UpdateProduct(ctx *fiber.Ctx, product *entities.ProductRegPayload, productId string) (pgconn.CommandTag, error)
	DeleteProduct(ctx *fiber.Ctx, productId string) (pgconn.CommandTag, error)
}

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) ProductRepo {
	return &productRepo{db}
}

func (r *productRepo) SearchProduct(ctx *fiber.Ctx, filter entities.FilterGetProducts) ([]entities.ProductList, error) {
	var products []entities.ProductList
	var createdAt time.Time

	query := "SELECT id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at FROM products"

	query += productConstructWhereQuery(filter)

	query += productConstructSortByQuery(filter.Price, filter.CreatedAt)

	query += " limit $1 offset $2"

	rows, err := r.db.Query(ctx.Context(), query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := entities.ProductList{}
		err := rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Notes, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &createdAt)
		if err != nil {
			return nil, err
		}

		product.CreatedAt = createdAt.Format(time.RFC3339)
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepo) SearchProductCustomer(ctx *fiber.Ctx, filter entities.FilterSku) ([]entities.CustomerProductList, error) {
	var products []entities.CustomerProductList
	var createdAt time.Time

	query := "SELECT id, name, sku, category, image_url, price, stock, location, created_at FROM products"

	query += custProductConstructWhereQuery(filter)

	query += custProductConstructSortByQuery(filter.Price)

	query += " limit $1 offset $2"

	rows, err := r.db.Query(ctx.Context(), query, filter.Limit, filter.Offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := entities.CustomerProductList{}
		err := rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Price, &product.Stock, &product.Location, &createdAt)
		if err != nil {
			return nil, err
		}
		product.CreatedAt = createdAt.Format(time.RFC3339)
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepo) AddProduct(ctx *fiber.Ctx, product *entities.ProductRegPayload) (string, time.Time, error) {
	var id string
	var createdAt time.Time

	// isAvail, err := strconv.ParseBool(product.IsAvailable)
	// if err != nil {
	// 	return "", time.Time{}, err
	// }

	statement := "INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, created_at"

	// Use QueryRow for inserting and getting the id back
	row := r.db.QueryRow(ctx.Context(), statement, product.Name, product.Sku, product.Category, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, *product.IsAvailable)
	if err := row.Scan(&id, &createdAt); err != nil {
		return "", time.Time{}, err
	}

	return id, createdAt, nil
}

func (r *productRepo) GetProductById(ctx *fiber.Ctx, id string) (*entities.Product, error) {
	var product entities.Product
	query := "SELECT * FROM products WHERE id = $1"

	// Use QueryRow to get a single row
	row := r.db.QueryRow(ctx.Context(), query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Notes, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt) // Add other fields as necessary
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) UpdateProduct(ctx *fiber.Ctx, product *entities.ProductRegPayload, productId string) (pgconn.CommandTag, error) {
	statement := "UPDATE products SET name = $1, sku = $2, category = $3, image_url = $4, notes = $5, price = $6, stock = $7, location = $8, is_available = $9 WHERE id = $10"
	// isAvail, err := strconv.ParseBool(product.IsAvailable)
	// if err != nil {
	// 	return pgconn.CommandTag{}, err
	// }

	res, err := r.db.Exec(ctx.Context(), statement, product.Name, product.Sku, product.Category, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, *product.IsAvailable, productId)

	return res, err
}

func (r *productRepo) DeleteProduct(ctx *fiber.Ctx, productId string) (pgconn.CommandTag, error) {
	statement := "DELETE FROM products WHERE id = $1"

	res, err := r.db.Exec(ctx.Context(), statement, productId)
	return res, err
}

func productConstructWhereQuery(filter entities.FilterGetProducts) string {
	whereSQL := []string{}

	if filter.Id != "" {
		whereSQL = append(whereSQL, " id = '"+filter.Id+"'")
	}

	if filter.Name != "" {
		whereSQL = append(whereSQL, " name ILIKE '%"+filter.Name+"%'")
	}

	if filter.IsAvailable == "true" {
		whereSQL = append(whereSQL, " is_available = '"+"1"+"'")
	}

	if filter.IsAvailable == "false" {
		whereSQL = append(whereSQL, " is_available = '"+"0"+"'")
	}

	if filter.Category != "" {
		whereSQL = append(whereSQL, " category = '"+filter.Category+"'")
	}

	if filter.Sku != "" {
		whereSQL = append(whereSQL, " sku = '"+filter.Sku+"'")
	}

	if filter.InStock == "true" {
		whereSQL = append(whereSQL, " stock > '"+"0"+"'")
	}

	if filter.InStock == "false" {
		whereSQL = append(whereSQL, " stock = '"+"0"+"'")
	}

	if len(whereSQL) > 0 {
		return " WHERE " + strings.Join(whereSQL, " AND ")
	}

	return ""
}

func productConstructSortByQuery(price string, createdAt string) string {
	sortBySQL := []string{}

	if price == "asc" {
		sortBySQL = append(sortBySQL, "price ASC")
	}

	if price == "desc" {
		sortBySQL = append(sortBySQL, "price DESC")
	}

	if createdAt == "asc" {
		sortBySQL = append(sortBySQL, "created_at ASC")
	}

	if createdAt == "desc" {
		sortBySQL = append(sortBySQL, "created_at DESC")
	}

	if len(sortBySQL) > 0 {
		return " ORDER BY " + strings.Join(sortBySQL, ", ")
	}

	return " ORDER BY created_at DESC"
}

func custProductConstructWhereQuery(filter entities.FilterSku) string {
	whereSQL := []string{}

	whereSQL = append(whereSQL, " is_available = TRUE")

	if filter.Name != "" {
		whereSQL = append(whereSQL, " name ILIKE '%"+filter.Name+"%'")
	}

	if filter.Category != "" {
		whereSQL = append(whereSQL, " category = '"+filter.Category+"'")
	}

	if filter.Sku != "" {
		whereSQL = append(whereSQL, " sku = '"+filter.Sku+"'")
	}

	if filter.InStock == "true" {
		whereSQL = append(whereSQL, " stock > '"+"0"+"'")
	}

	if filter.InStock == "false" {
		whereSQL = append(whereSQL, " stock = '"+"0"+"'")
	}

	if len(whereSQL) > 0 {
		return " WHERE " + strings.Join(whereSQL, " AND ")
	}

	return ""
}

func custProductConstructSortByQuery(price string) string {
	sortBySQL := []string{}

	if price == "asc" {
		sortBySQL = append(sortBySQL, "price ASC")
	}

	if price == "desc" {
		sortBySQL = append(sortBySQL, "price DESC")
	}

	if len(sortBySQL) > 0 {
		return " ORDER BY " + strings.Join(sortBySQL, ", ")
	}

	return " ORDER BY created_at DESC"
}
