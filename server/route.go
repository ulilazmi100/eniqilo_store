package server

import (
	"eniqilo_store/controller"
	"eniqilo_store/middleware"
	"eniqilo_store/repo"
	"eniqilo_store/svc"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *Server) RegisterRoute() {
	mainRoute := s.app.Group("/v1")

	registerUserRoute(mainRoute, s.dbPool)
	registerCustomerRoute(mainRoute, s.dbPool)
	registerProductRoute(mainRoute, s.dbPool)
}

func registerUserRoute(r fiber.Router, db *pgxpool.Pool) {
	ctr := controller.NewUserController(svc.NewUserSvc(repo.NewUserRepo(db)))
	userGroup := r.Group("/staff")

	newRoute(userGroup, "POST", "/register", ctr.Register)
	newRoute(userGroup, "POST", "/login", ctr.Login)
}

func registerProductRoute(r fiber.Router, db *pgxpool.Pool) {
	ctr := controller.NewProductController(svc.NewProductSvc(repo.NewProductRepo(db)))
	productGroup := r.Group("/product")

	newRouteWithAuth(productGroup, "POST", "/", ctr.Register)
	newRouteWithAuth(productGroup, "GET", "/", ctr.Search)
	newRouteWithAuth(productGroup, "PUT", "/:productId", ctr.Update)
	newRouteWithAuth(productGroup, "DELETE", "/:productId", ctr.Delete)

	newRoute(productGroup, "GET", "/customer", ctr.SearchSKU)

	productCheckoutGroup := r.Group("/product/checkout")
	ctrTransaction := controller.NewTransactionController(svc.NewTransactionSvc(repo.NewTransactionRepo(db)))

	newRouteWithAuth(productCheckoutGroup, "POST", "/", ctrTransaction.Checkout)
	newRouteWithAuth(productCheckoutGroup, "GET", "/history", ctrTransaction.History)
}

func registerCustomerRoute(r fiber.Router, db *pgxpool.Pool) {
	ctr := controller.NewCustomerController(svc.NewCustomerSvc(repo.NewCustomerRepo(db)))
	customerGroup := r.Group("/customer")

	newRouteWithAuth(customerGroup, "POST", "/register", ctr.Register)
	newRouteWithAuth(customerGroup, "GET", "/", ctr.Search)
}

func newRoute(router fiber.Router, method, path string, handler fiber.Handler) {
	router.Add(method, path, handler)
}

func newRouteWithAuth(router fiber.Router, method, path string, handler fiber.Handler) {
	router.Add(method, path, middleware.Auth, handler)
}
