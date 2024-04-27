package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"rent_gpu_be/internal/delivery/http"
	"rent_gpu_be/internal/delivery/http/route"
	"rent_gpu_be/internal/repository"
	"rent_gpu_be/internal/usecase"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	gpuRepository := repository.NewGpuRepository(config.Log)
	orderRepository := repository.NewOrderRepository(config.Log)
	userRepository := repository.NewUseriRepository(config.Log)

	// setup use cases
	gpuUseCase := usecase.NewGpuUseCase(config.DB, config.Log, config.Validate, gpuRepository)
	orderUseCase := usecase.NewOrderUseCase(config.DB, config.Log, config.Validate, orderRepository)
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)

	// setup controller
	gpuController := http.NewGpuController(gpuUseCase, config.Log)
	orderController := http.NewOrderController(orderUseCase, config.Log)
	userController := http.NewUserController(userUseCase, config.Log)

	// setup middleware
	//authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:             config.App,
		GpuController:   gpuController,
		OrderController: orderController,
		UserController:  userController,
		//AuthMiddleware:    authMiddleware,
	}
	routeConfig.Setup()
}
