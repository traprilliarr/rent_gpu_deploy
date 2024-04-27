package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"rent_gpu_be/internal/model"
	"rent_gpu_be/internal/usecase"
)

type OrderController struct {
	UseCase *usecase.OrderUseCase
	Log     *logrus.Logger
}

func NewOrderController(useCase *usecase.OrderUseCase, log *logrus.Logger) *OrderController {
	return &OrderController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {

	var requestBody model.OrderRequest

	// Parse the JSON body into the defined struct
	if err := ctx.BodyParser(&requestBody); err != nil {
		// If there's an error parsing the body, return an error response
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	responses, err := c.UseCase.CreateOrder(ctx.UserContext(), requestBody)
	if err != nil {
		c.Log.WithError(err).Error("error create order")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[*model.OrderResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}

func (c *OrderController) Update(ctx *fiber.Ctx) error {

	var requestBody model.OrderUpdatedRequest

	// Parse the JSON body into the defined struct
	if err := ctx.BodyParser(&requestBody); err != nil {
		// If there's an error parsing the body, return an error response
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	responses, err := c.UseCase.UpdateOrder(ctx.UserContext(), requestBody)
	if err != nil {
		c.Log.WithError(err).Error("error updated order")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[*model.OrderResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}

func (c *OrderController) GetOrderByUserID(ctx *fiber.Ctx) error {

	var requestBody model.OrderUserRequest

	// Parse the JSON body into the defined struct
	if err := ctx.BodyParser(&requestBody); err != nil {
		// If there's an error parsing the body, return an error response
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	responses, err := c.UseCase.GetAllOrderByUserID(ctx.UserContext(), requestBody)
	if err != nil {
		c.Log.WithError(err).Error("error updated order")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[[]*model.OrderResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}

func (c *OrderController) GetOrderByID(ctx *fiber.Ctx) error {

	var requestBody model.OrderUpdatedRequest

	// Parse the JSON body into the defined struct
	if err := ctx.BodyParser(&requestBody); err != nil {
		// If there's an error parsing the body, return an error response
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	responses, err := c.UseCase.GetOrderByID(ctx.UserContext(), requestBody)
	if err != nil {
		c.Log.WithError(err).Error("error updated order")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[*model.OrderResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}
