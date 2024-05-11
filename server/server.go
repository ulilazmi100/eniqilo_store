package server

import (
	"eniqilo_store/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"

	configs "eniqilo_store/cfg"
)

type Server struct {
	dbPool    *pgxpool.Pool
	app       *fiber.App
	validator *validator.Validate
}

func NewServer(db *pgxpool.Pool) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	validate := validator.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	return &Server{
		dbPool:    db,
		app:       app,
		validator: validate,
	}
}

func (s *Server) Run(config configs.Config) error {
	return s.app.Listen(":" + config.APPPort)
}
