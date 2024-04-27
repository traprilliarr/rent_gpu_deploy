package route

import (
	"github.com/gofiber/fiber/v2"
	"rent_gpu_be/internal/delivery/http"
)

type RouteConfig struct {
	App             *fiber.App
	GpuController   *http.GpuController
	OrderController *http.OrderController
	UserController  *http.UserController
	AuthMiddleware  fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	//c.App.Post("/api/users", c.UserController.Register)
	//c.App.Post("/api/users/_login", c.UserController.Login)
	c.App.Get("/api/gpu/list", c.GpuController.List)
	c.App.Get("/api/gpu/list/:id", c.GpuController.ById)
	c.App.Post("/api/order/create", c.OrderController.Create)
	c.App.Put("/api/order/update", c.OrderController.Update)
	c.App.Get("/api/order/get", c.OrderController.GetOrderByUserID)
	c.App.Get("/api/order/detail", c.OrderController.GetOrderByID)
	c.App.Get("/api/auth/nonce", c.UserController.GetNonce)
	c.App.Get("/api/auth/login", c.UserController.Login)

}

func (c *RouteConfig) SetupAuthRoute() {
	//c.App.Use(c.AuthMiddleware)
	//c.App.Get("/api/contacts/:contactId", c.ContactController.Get)
}
