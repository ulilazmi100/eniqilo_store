package svc

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/repo"
	"eniqilo_store/responses"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type TransactionSvc interface {
	Checkout(ctx *fiber.Ctx, newTransaction entities.TransactionPayload) error
	SearchTransaction(ctx *fiber.Ctx, transaction entities.FilterGetTransactions) ([]entities.Transaction, error)
}

type transactionSvc struct {
	repo  repo.TransactionRepo
	mutex sync.Mutex
}

func NewTransactionSvc(repo repo.TransactionRepo) TransactionSvc {
	return &transactionSvc{repo: repo}
}

func (s *transactionSvc) Checkout(ctx *fiber.Ctx, newTransaction entities.TransactionPayload) error {
	if err := newTransaction.Validate(); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	var totalPrice int
	var ownedProductDetails []entities.ProductDetail

	if newTransaction.CustomerId != "" {
		_, err := s.repo.GetCustomerById(ctx, newTransaction.CustomerId)
		if err == pgx.ErrNoRows {
			return responses.NewNotFoundError("customerId not found")
		}
	}

	if len(newTransaction.ProductDetails) == 0 {
		return responses.NewBadRequestError("productDetails is empty")
	}

	for _, product_detail := range newTransaction.ProductDetails {
		if product_detail.Quantity < 1 {
			return responses.NewBadRequestError("One of product quantity is below 1")
		}
		if product_detail.ProductId == "" {
			return responses.NewBadRequestError("Empty ProductId")
		}
		price, stock, is_avail, err := s.repo.GetProductById(ctx, product_detail.ProductId)
		if err == pgx.ErrNoRows {
			return responses.NewNotFoundError("one of productIds is not found")
		} else if err != nil {
			return responses.NewNotFoundError(err.Error())
		}
		if !is_avail {
			return responses.NewBadRequestError("one of productIds isAvailable == false")
		}
		if stock < product_detail.Quantity {
			return responses.NewBadRequestError("one of productIds stock is not enough")
		}
		totalPrice += price * product_detail.Quantity

		ownedProductDetails = append(ownedProductDetails, entities.ProductDetail{ProductId: product_detail.ProductId, Quantity: stock - product_detail.Quantity})
	}

	if newTransaction.Paid < totalPrice {
		return responses.NewBadRequestError("paid is not enough based on all bought product")
	}

	if (newTransaction.Paid - totalPrice) != newTransaction.Change {
		return responses.NewBadRequestError("change is not right, based on all bought product, and what is paid")
	}

	err := s.repo.AddTransaction(ctx, &newTransaction)
	if err != nil {
		return responses.NewInternalServerError(err.Error())
	}

	for _, product_detail := range ownedProductDetails {
		_, err := s.repo.UpdateProduct(ctx, product_detail.Quantity, product_detail.ProductId)
		if err != nil {
			return responses.NewInternalServerError(err.Error())
		}
	}

	return nil
}

func (s *transactionSvc) SearchTransaction(ctx *fiber.Ctx, transaction entities.FilterGetTransactions) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	transactions, err := s.repo.SearchTransaction(ctx, transaction)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entities.Transaction{}, nil
		}
		return nil, err
	}

	return transactions, nil
}
