package svc

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/repo"
	"eniqilo_store/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type CustomerSvc interface {
	Register(ctx *fiber.Ctx, newCustomer entities.CustomerRegPayload) (string, error)
	SearchCustomer(ctx *fiber.Ctx, customer entities.CustomerFilter) ([]entities.CustomerList, error)
}

type customerSvc struct {
	repo repo.CustomerRepo
}

func NewCustomerSvc(repo repo.CustomerRepo) CustomerSvc {
	return &customerSvc{repo}
}

func (s *customerSvc) Register(ctx *fiber.Ctx, newCustomer entities.CustomerRegPayload) (string, error) {
	customer := entities.NewCustomer(newCustomer.PhoneNumber, newCustomer.Name)

	if err := customer.Validate(); err != nil {
		return "", responses.NewBadRequestError(err.Error())
	}

	existingCustomer, err := s.repo.GetCustomerByPhone(ctx, newCustomer.PhoneNumber)
	if err != nil {
		if err != pgx.ErrNoRows {
			return "", err
		}
	}

	if existingCustomer != nil {
		return "", responses.NewConflictError("PhoneNumber already exist")
	}

	id, err := s.repo.CreateCustomer(ctx, &newCustomer)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *customerSvc) SearchCustomer(ctx *fiber.Ctx, customer entities.CustomerFilter) ([]entities.CustomerList, error) {
	var customers []entities.CustomerList

	customers, err := s.repo.SearchCustomer(ctx, customer.PhoneNumber, customer.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entities.CustomerList{}, nil
		}
		return nil, err
	}

	return customers, nil
}
