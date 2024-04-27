package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rent_gpu_be/internal/converter"
	"rent_gpu_be/internal/entity"
	"rent_gpu_be/internal/model"
	"rent_gpu_be/internal/repository"
)

type GpuUseCase struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	Validate      *validator.Validate
	GpuRepository *repository.GpuRepository
}

func NewGpuUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	gpuRepository *repository.GpuRepository) *GpuUseCase {
	return &GpuUseCase{
		DB:            db,
		Log:           logger,
		Validate:      validate,
		GpuRepository: gpuRepository,
	}
}

func (c *GpuUseCase) GetAll(ctx context.Context) ([]*model.GpuResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	gpu := new([]entity.Gpu)
	if err := c.GpuRepository.FindAll(tx, gpu); err != nil {
		c.Log.WithError(err).Error("error getting cpu")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting cpu")
		return nil, fiber.ErrInternalServerError
	}

	return converter.GpuToResponses(gpu), nil
}

func (c *GpuUseCase) GetByID(ctx context.Context, request string) (*model.GpuResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	gpu := new(entity.Gpu)
	if err := c.GpuRepository.FindById(tx, gpu, request); err != nil {
		c.Log.WithError(err).Error("error getting cpu")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting cpu")
		return nil, fiber.ErrInternalServerError
	}

	return converter.GpuToResponse(gpu), nil
}
