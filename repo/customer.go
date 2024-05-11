package repo

import (
	"eniqilo_store/db/entities"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepo interface {
	SearchCustomer(ctx *fiber.Ctx, phone string, name string) ([]entities.CustomerList, error)
	CreateCustomer(ctx *fiber.Ctx, customer *entities.CustomerRegPayload) (string, error)
	GetCustomerByPhone(ctx *fiber.Ctx, phone string) (*entities.Customer, error)
}

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomerRepo(db *pgxpool.Pool) CustomerRepo {
	return &customerRepo{db}
}

func (r *customerRepo) SearchCustomer(ctx *fiber.Ctx, phone string, name string) ([]entities.CustomerList, error) {
	var customers []entities.CustomerList
	query := "SELECT id, name, phone FROM customers"

	query += custConstructWhereQuery(phone, name)

	rows, err := r.db.Query(ctx.Context(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		customer := entities.CustomerList{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *customerRepo) CreateCustomer(ctx *fiber.Ctx, customer *entities.CustomerRegPayload) (string, error) {
	var id string
	statement := "INSERT INTO customers (name, phone) VALUES ($1, $2) RETURNING id"

	// Use QueryRow for inserting and getting the id back
	row := r.db.QueryRow(ctx.Context(), statement, customer.Name, customer.PhoneNumber)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *customerRepo) GetCustomerByPhone(ctx *fiber.Ctx, phone string) (*entities.Customer, error) {
	var customer entities.Customer
	query := "SELECT * FROM customers WHERE phone = $1"

	// Use QueryRow to get a single row
	row := r.db.QueryRow(ctx.Context(), query, phone)
	err := row.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt) // Add other fields as necessary
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func custConstructWhereQuery(phone string, name string) string {
	whereSQL := []string{}

	if phone != "" {
		whereSQL = append(whereSQL, " phone ILIKE '"+phone+"%'")
	}

	if name != "" {
		whereSQL = append(whereSQL, " name ILIKE '%"+name+"%'")
	}

	if len(whereSQL) > 0 {
		return " WHERE " + strings.Join(whereSQL, " AND ")
	}

	return ""
}
