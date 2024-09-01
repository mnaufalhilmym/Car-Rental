package config

import (
	"carrental/internal/controller"
	"carrental/internal/repository"
	"carrental/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	Validator *validator.Validate
	Router    *gin.Engine
}

func Bootstrap(conf BootstrapConfig) {
	// Repository
	customerRepository := repository.NewCustomerRepository(conf.DB)
	carRepository := repository.NewCarRepository(conf.DB)
	bookingRepository := repository.NewBookingRepository(conf.DB)

	// Usecase
	customerUsecase := usecase.NewCustomerUsecase(conf.DB, conf.Validator, customerRepository)
	carUsecase := usecase.NewCarUsecase(conf.DB, conf.Validator, carRepository)
	bookingUsecase := usecase.NewBookingUsecase(conf.DB, conf.Validator, bookingRepository, customerRepository, carRepository)

	// Controller
	customerController := controller.NewCustomerController(customerUsecase)
	carController := controller.NewCarController(carUsecase)
	bookingController := controller.NewBookingController(bookingUsecase)

	routeConfig := controller.RouteConfig{
		Router:             conf.Router,
		CustomerController: customerController,
		CarController:      carController,
		BookingController:  bookingController,
	}

	routeConfig.ConfigureRoutes()
}
