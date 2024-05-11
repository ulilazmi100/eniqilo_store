package middleware

import (
	"errors"
	"strings"

	"eniqilo_store/crypto"
	"eniqilo_store/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return responses.NewUnauthorizedError("token not found")
	}

	splitted := strings.Split(auth, " ")

	if splitted[0] != "Bearer" {
		return responses.NewUnauthorizedError("invalid token")
	}

	payload, err := crypto.VerifyToken(splitted[1])
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return responses.NewUnauthorizedError("token expired")
		}
		return responses.NewUnauthorizedError(err.Error())
	}

	ctx.Locals("user_id", payload.Id)
	ctx.Locals("phone", payload.Phone)
	ctx.Locals("name", payload.Name)
	return ctx.Next()
}
