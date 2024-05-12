package repo

import (
	"eniqilo_store/db/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo interface {
	GetUser(ctx *fiber.Ctx, phone string) (*entities.User, error)
	CreateUser(ctx *fiber.Ctx, user *entities.RegistrationPayload, hashPassword string) (string, error)
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) GetUser(ctx *fiber.Ctx, phone string) (*entities.User, error) {
	var user entities.User
	query := "SELECT * FROM users WHERE phone = $1"

	// Use QueryRow to get a single row
	row := r.db.QueryRow(ctx.Context(), query, phone)
	err := row.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Password, &user.CreatedAt, &user.UpdatedAt) // Add other fields as necessary
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CreateUser(ctx *fiber.Ctx, user *entities.RegistrationPayload, hashPassword string) (string, error) {
	var id string
	statement := "INSERT INTO users (name, phone, password) VALUES ($1, $2, $3) RETURNING id"

	// Use QueryRow for inserting and getting the id back
	row := r.db.QueryRow(ctx.Context(), statement, user.Name, user.PhoneNumber, hashPassword)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
