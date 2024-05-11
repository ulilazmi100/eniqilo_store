package controller

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/responses"
	"eniqilo_store/svc"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	svc svc.ProductSvc
}

func NewProductController(svc svc.ProductSvc) *ProductController {
	return &ProductController{svc: svc}
}

func (c *ProductController) Register(ctx *fiber.Ctx) error {
	var newProduct entities.ProductRegPayload
	if err := ctx.BodyParser(&newProduct); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	productId, createdAt, err := c.svc.Register(ctx, newProduct)
	if err != nil {
		return err
	}

	respData := fiber.Map{
		"id":        productId,
		"createdAt": createdAt.Format(time.RFC3339),
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "success",
		"data":    respData,
	})
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}

	var newProduct entities.ProductRegPayload
	if err := ctx.BodyParser(&newProduct); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	err = c.svc.Update(ctx, newProduct, id.String())
	if err != nil {
		return err
	}

	return ctx.Status(201).JSON(fiber.Map{})
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}

	err = c.svc.DeleteProduct(ctx, id.String())
	if err != nil {
		return err
	}

	return ctx.Status(201).JSON(fiber.Map{})
}

func (c *ProductController) Search(ctx *fiber.Ctx) error {
	var product entities.FilterGetProducts

	if err := ctx.QueryParser(&product); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	resp, err := c.svc.SearchProduct(ctx, product)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}

func (c *ProductController) SearchSKU(ctx *fiber.Ctx) error {
	var product entities.FilterSku

	if err := ctx.QueryParser(&product); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	resp, err := c.svc.SearchSKU(ctx, product)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}
