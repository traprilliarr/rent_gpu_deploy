package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"rent_gpu_be/internal/model"
	"rent_gpu_be/internal/usecase"
)

type UserController struct {
	UseCase *usecase.UserUseCase
	Log     *logrus.Logger
}

func NewUserController(useCase *usecase.UserUseCase, log *logrus.Logger) *UserController {
	return &UserController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *UserController) GetNonce(ctx *fiber.Ctx) error {

	responses, err := c.UseCase.GetNonce(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("error updated order")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[*model.NonceResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {

	var requestBody model.AuthRequest

	// Parse the JSON body into the defined struct
	if err := ctx.BodyParser(&requestBody); err != nil {
		// If there's an error parsing the body, return an error response
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	responses, err := c.UseCase.Login(ctx.UserContext(), requestBody)
	if err != nil {
		c.Log.WithError(err).Error("error updated order")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[string]{
		HttpCode: 200,
		Data:     responses,
	})
}
