package svc

import (
	"eniqilo_store/crypto"
	"eniqilo_store/db/entities"
	"eniqilo_store/repo"
	"eniqilo_store/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type UserSvc interface {
	Register(ctx *fiber.Ctx, newUser entities.RegistrationPayload) (string, string, error)
	Login(ctx *fiber.Ctx, user entities.Credential) (string, string, string, error)
}

type userSvc struct {
	repo repo.UserRepo
}

func NewUserSvc(repo repo.UserRepo) UserSvc {
	return &userSvc{repo}
}

func (s *userSvc) Register(ctx *fiber.Ctx, newUser entities.RegistrationPayload) (string, string, error) {
	user := entities.NewUser(newUser.PhoneNumber, newUser.Name, newUser.Password)

	if err := user.Validate(); err != nil {
		return "", "", responses.NewBadRequestError(err.Error())
	}

	existingUser, err := s.repo.GetUser(ctx, newUser.PhoneNumber)
	if err != nil {
		if err != pgx.ErrNoRows {
			return "", "", err
		}
	}

	if existingUser != nil {
		return "", "", responses.NewConflictError("user already exist")
	}

	hashedPassword, err := crypto.GenerateHashedPassword(newUser.Password)
	if err != nil {
		return "", "", err
	}

	id, err := s.repo.CreateUser(ctx, &newUser, hashedPassword)
	if err != nil {
		return "", "", err
	}

	token, err := crypto.GenerateToken(id, newUser.PhoneNumber, newUser.Name)
	if err != nil {
		return "", "", err
	}

	return id, token, nil
}

func (s *userSvc) Login(ctx *fiber.Ctx, creds entities.Credential) (string, string, string, error) {
	if err := creds.Validate(); err != nil {
		return "", "", "", responses.NewBadRequestError(err.Error())
	}

	user, err := s.repo.GetUser(ctx, creds.PhoneNumber)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", "", "", responses.NewNotFoundError("user not found")
		}
		return "", "", "", err
	}

	err = crypto.VerifyPassword(creds.Password, user.Password)
	if err != nil {
		return "", "", "", responses.NewBadRequestError("wrong password!")
	}

	token, err := crypto.GenerateToken(user.Id, user.PhoneNumber, user.Name)
	if err != nil {
		return "", "", "", responses.NewBadRequestError(err.Error())
	}

	return user.Id, user.Name, token, nil
}
