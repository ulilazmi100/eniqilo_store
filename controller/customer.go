package controller

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/responses"
	"eniqilo_store/svc"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	svc svc.CustomerSvc
}

func NewCustomerController(svc svc.CustomerSvc) *CustomerController {
	return &CustomerController{svc: svc}
}

func (c *CustomerController) Register(ctx *fiber.Ctx) error {
	var newCustomer entities.CustomerRegPayload
	if err := ctx.BodyParser(&newCustomer); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	customerId, err := c.svc.Register(ctx, newCustomer)
	if err != nil {
		return err
	}

	respData := fiber.Map{
		"customerId": customerId,
		"phone":      newCustomer.PhoneNumber,
		"name":       newCustomer.Name,
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    respData,
	})
}

func (c *CustomerController) Search(ctx *fiber.Ctx) error {
	var customer entities.CustomerFilter

	if err := ctx.QueryParser(&customer); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	if customer.PhoneNumber != "" {
		customer.PhoneNumber = "+" + customer.PhoneNumber
	}

	resp, err := c.svc.SearchCustomer(ctx, customer)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    resp,
	})
}
