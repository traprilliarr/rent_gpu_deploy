package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rent_gpu_be/internal/converter"
	"rent_gpu_be/internal/entity"
	"rent_gpu_be/internal/model"
	"rent_gpu_be/internal/repository"
	"time"
)

type OrderUseCase struct {
	DB              *gorm.DB
	Log             *logrus.Logger
	Validate        *validator.Validate
	OrderRepository *repository.OrderRepository
}

func NewOrderUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	orderRepository *repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		DB:              db,
		Log:             logger,
		Validate:        validate,
		OrderRepository: orderRepository,
	}
}

func (c *OrderUseCase) CreateOrder(ctx context.Context, request model.OrderRequest) (*model.OrderResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Validate the OrderRequest
	if err := c.Validate.Struct(request); err != nil {
		// If validation fails, print the validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			c.Log.WithError(err).Error("validation error")
		}
		return nil, err
	}

	order := new(entity.Order)

	order.ID = uuid.NewString()
	order.Email = request.Email
	order.Telegram = request.Telegram
	order.PaymentAddress = "0x123hb1321n3mnk128218128921981"
	order.Value = request.Value
	order.Hash = request.Hash
	order.Status = "Checking Payment"
	order.GpuID = request.GpuID
	order.UserID = request.UserID
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	if err := c.OrderRepository.Create(tx, order); err != nil {
		c.Log.WithError(err).Error("error create order")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting order")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}

func (c *OrderUseCase) UpdateOrder(ctx context.Context, request model.OrderUpdatedRequest) (*model.OrderResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Validate the OrderRequest
	if err := c.Validate.Struct(request); err != nil {
		// If validation fails, print the validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			c.Log.WithError(err).Error("validation error")
		}
		return nil, err
	}

	order := new(entity.Order)

	err := c.OrderRepository.FindById(tx, order, request.OrderID)
	if err != nil {
		c.Log.WithError(err).Error("order not found with id : " + request.OrderID)
		return nil, fiber.ErrNotFound
	}

	if order.Status == "Success" {
		c.Log.WithError(err).Error("order with this id already confirmed")
		return nil, errors.New("order with this id already confirmed")
	}
	order.Status = "Success"
	order.UpdatedAt = time.Now()

	if err := c.OrderRepository.Update(tx, order); err != nil {
		c.Log.WithError(err).Error("error updated order")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting order")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}

func (c *OrderUseCase) GetAllOrderByUserID(ctx context.Context, request model.OrderUserRequest) ([]*model.OrderResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		// If validation fails, print the validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			c.Log.WithError(err).Error("validation error")
		}
		return nil, err
	}

	gpu := new([]entity.Order)

	if err := c.OrderRepository.FindAllByUserId(tx, gpu, request.UserID); err != nil {
		c.Log.WithError(err).Error("error getting order")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting order")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponses(gpu), nil
}

func (c *OrderUseCase) GetOrderByID(ctx context.Context, request model.OrderUpdatedRequest) (*model.OrderResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		// If validation fails, print the validation errors
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			c.Log.WithError(err).Error("validation error")
		}
		return nil, err
	}

	gpu := new(entity.Order)

	if err := c.OrderRepository.FindById(tx, gpu, request.OrderID); err != nil {
		c.Log.WithError(err).Error("error getting order")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting order")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(gpu), nil
}
