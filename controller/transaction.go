package controller

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/responses"
	"eniqilo_store/svc"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	svc svc.TransactionSvc
}

func NewTransactionController(svc svc.TransactionSvc) *TransactionController {
	return &TransactionController{svc: svc}
}

func (c *TransactionController) Checkout(ctx *fiber.Ctx) error {
	var newTransaction entities.TransactionPayload
	if err := ctx.BodyParser(&newTransaction); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	validate := validator.New()
	err := validate.Struct(newTransaction)
	if err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	err = c.svc.Checkout(ctx, newTransaction)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{})
}

func (c *TransactionController) History(ctx *fiber.Ctx) error {
	var transaction entities.FilterGetTransactions

	if err := ctx.QueryParser(&transaction); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	if transaction.Limit == 0 {
		transaction.Limit = 5
	}

	if transaction.Limit < 0 || transaction.Offset < 0 {
		return responses.NewBadRequestError("invalid query param")
	}

	resp, err := c.svc.SearchTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}
