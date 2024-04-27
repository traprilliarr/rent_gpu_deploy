package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"rent_gpu_be/internal/model"
	"rent_gpu_be/internal/usecase"
)

type GpuController struct {
	UseCase *usecase.GpuUseCase
	Log     *logrus.Logger
}

func NewGpuController(useCase *usecase.GpuUseCase, log *logrus.Logger) *GpuController {
	return &GpuController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *GpuController) List(ctx *fiber.Ctx) error {

	responses, err := c.UseCase.GetAll(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("error get all cpu")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[[]*model.GpuResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}

func (c *GpuController) ById(ctx *fiber.Ctx) error {

	query := ctx.Params("id")

	responses, err := c.UseCase.GetByID(ctx.UserContext(), query)
	if err != nil {
		c.Log.WithError(err).Error("error get all cpu")
		// Return error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[error]{
			HttpCode: fiber.StatusInternalServerError,
			Errors:   err.Error(),
		})
	}
	return ctx.JSON(model.WebResponse[*model.GpuResponse]{
		HttpCode: 200,
		Data:     responses,
	})
}
