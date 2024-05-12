package controller

import (
	"eniqilo_store/db/entities"
	"eniqilo_store/responses"
	"eniqilo_store/svc"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	svc svc.UserSvc
}

func NewUserController(svc svc.UserSvc) *UserController {
	return &UserController{svc: svc}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var newUser entities.RegistrationPayload
	if err := ctx.BodyParser(&newUser); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	userId, accessToken, err := c.svc.Register(ctx, newUser)
	if err != nil {
		return err
	}

	respData := fiber.Map{
		"userId":      userId,
		"phoneNumber": newUser.PhoneNumber,
		"name":        newUser.Name,
		"accessToken": accessToken,
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    respData,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var user entities.RegistrationPayload
	if err := ctx.BodyParser(&user); err != nil {
		return responses.NewBadRequestError(err.Error())
	}

	loginPayload := entities.Credential{
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}

	userId, name, accessToken, err := c.svc.Login(ctx, loginPayload)
	if err != nil {
		return err
	}

	respData := fiber.Map{
		"userId":      userId,
		"phoneNumber": user.PhoneNumber,
		"name":        name,
		"accessToken": accessToken,
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    respData,
	})
}
