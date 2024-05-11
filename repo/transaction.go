package repo

import (
	"eniqilo_store/db/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepo interface {
	SearchTransaction(ctx *fiber.Ctx, filter entities.FilterGetTransactions) ([]entities.Transaction, error)
	AddTransaction(ctx *fiber.Ctx, transaction *entities.TransactionPayload) error
	GetProductById(ctx *fiber.Ctx, id string) (int, int, bool, error)
	UpdateProduct(ctx *fiber.Ctx, stock int, productId string) (pgconn.CommandTag, error)
	GetCustomerById(ctx *fiber.Ctx, id string) (*entities.Customer, error)
}

type transactionRepo struct {
	db *pgxpool.Pool
}

func NewTransactionRepo(db *pgxpool.Pool) TransactionRepo {
	return &transactionRepo{db}
}

func (r *transactionRepo) SearchTransaction(ctx *fiber.Ctx, filter entities.FilterGetTransactions) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	query := "SELECT id, customer_id, product_details, paid, change, created_at FROM transactions"

	if filter.CustomerId != "" {
		query += " WHERE customer_id = '" + filter.CustomerId + "'"
	}
	if filter.CreatedAt != "" {
		if filter.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if filter.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}

	rows, err := r.db.Query(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		transaction := entities.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.ProductDetails, transaction.Paid, transaction.Change, transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *transactionRepo) AddTransaction(ctx *fiber.Ctx, transaction *entities.TransactionPayload) error {
	statement := "INSERT INTO transactions (customer_id, product_details, paid, change) VALUES ($1, $2, $3, $4)"

	// Use QueryRow for inserting and getting the id back
	_, err := r.db.Exec(ctx.Context(), statement, transaction.CustomerId, transaction.ProductDetails, transaction.Paid, transaction.Change)
	if err != nil {
		return err
	}

	return err
}

func (r *transactionRepo) GetProductById(ctx *fiber.Ctx, id string) (int, int, bool, error) {
	var price int
	var stock int
	var is_avail bool
	query := "SELECT price, stock, is_avail FROM products WHERE id = $1"

	// Use QueryRow to get a single row
	row := r.db.QueryRow(ctx.Context(), query, id)
	err := row.Scan(&price, &stock, &is_avail) // Add other fields as necessary
	if err != nil {
		return 0, 0, is_avail, err
	}

	return price, stock, is_avail, nil
}

func (r *transactionRepo) UpdateProduct(ctx *fiber.Ctx, stock int, productId string) (pgconn.CommandTag, error) {
	statement := "UPDATE products SET stock = $1 WHERE id = $2"

	res, err := r.db.Exec(ctx.Context(), statement, stock, productId)

	return res, err
}

func (r *transactionRepo) GetCustomerById(ctx *fiber.Ctx, id string) (*entities.Customer, error) {
	var customer entities.Customer
	query := "SELECT * FROM customers WHERE id = $1"

	// Use QueryRow to get a single row
	row := r.db.QueryRow(ctx.Context(), query, id)
	err := row.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt) // Add other fields as necessary
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
