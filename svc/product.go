package svc

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/repo"
	"eniqilo_store/responses"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type ProductSvc interface {
	Register(ctx *fiber.Ctx, newProduct entities.ProductRegPayload) (string, time.Time, error)
	Update(ctx *fiber.Ctx, newProduct entities.ProductRegPayload, productID string) error
	SearchProduct(ctx *fiber.Ctx, product entities.FilterGetProducts) ([]entities.ProductList, error)
	SearchSKU(ctx *fiber.Ctx, product entities.FilterSku) ([]entities.CustomerProductList, error)
	DeleteProduct(ctx *fiber.Ctx, productId string) error
}

type productSvc struct {
	repo repo.ProductRepo
}

func NewProductSvc(repo repo.ProductRepo) ProductSvc {
	return &productSvc{repo}
}

func (s *productSvc) Register(ctx *fiber.Ctx, newProduct entities.ProductRegPayload) (string, time.Time, error) {
	if err := newProduct.Validate(); err != nil {
		return "", time.Time{}, responses.NewBadRequestError(err.Error())
	}

	id, time, err := s.repo.AddProduct(ctx, &newProduct)
	if err != nil {
		return "", time, err
	}

	return id, time, nil
}

func (s *productSvc) Update(ctx *fiber.Ctx, newProduct entities.ProductRegPayload, productID string) error {
	if err := newProduct.Validate(); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	res, err := s.repo.UpdateProduct(ctx, &newProduct, productID)
	if res.RowsAffected() == 0 {
		return responses.NewNotFoundError(err.Error())
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *productSvc) SearchProduct(ctx *fiber.Ctx, product entities.FilterGetProducts) ([]entities.ProductList, error) {
	var products []entities.ProductList

	product.Category = entities.CategoryChecker(product.Category)

	products, err := s.repo.SearchProduct(ctx, product)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entities.ProductList{}, nil
		}
		return nil, err
	}

	return products, nil
}

func (s *productSvc) SearchSKU(ctx *fiber.Ctx, product entities.FilterSku) ([]entities.CustomerProductList, error) {
	var products []entities.CustomerProductList

	product.Category = entities.CategoryChecker(product.Category)

	products, err := s.repo.SearchProductCustomer(ctx, product)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []entities.CustomerProductList{}, nil
		}
		return nil, err
	}

	return products, nil
}

func (s *productSvc) DeleteProduct(ctx *fiber.Ctx, productId string) error {

	res, err := s.repo.DeleteProduct(ctx, productId)

	if res.RowsAffected() == 0 {
		return responses.NewNotFoundError(err.Error())
	}

	if err != nil {
		return responses.NewInternalServerError(err.Error())
	}

	return err
}
