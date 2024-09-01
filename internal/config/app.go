package config

import (
	"carrental/internal/controller"
	"carrental/internal/repository"
	"carrental/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func Bootstrap(conf BootstrapConfig) {
	// Repository
	membershipRepository := repository.NewMembershipRepository(conf.DB)
	customerRepository := repository.NewCustomerRepository(conf.DB)
	carRepository := repository.NewCarRepository(conf.DB)
	bookingTypeRepository := repository.NewBookingTypeRepository(conf.DB)
	driverRepository := repository.NewDriverRepository(conf.DB)
	driverIncentiveRepository := repository.NewDriverIncentiveRepository(conf.DB)
	bookingRepository := repository.NewBookingRepository(conf.DB)

	// Usecase
	membershipUsecase := usecase.NewMembershipUsecase(conf.DB, membershipRepository)
	customerUsecase := usecase.NewCustomerUsecase(conf.DB, customerRepository, membershipRepository)
	carUsecase := usecase.NewCarUsecase(conf.DB, carRepository)
	bookingTypeUsecase := usecase.NewBookingTypeUsecase(conf.DB, bookingTypeRepository)
	driverUsecase := usecase.NewDriverUsecase(conf.DB, driverRepository)
	bookingUsecase := usecase.NewBookingUsecase(conf.DB, bookingRepository, customerRepository, carRepository, bookingTypeRepository, driverRepository, driverIncentiveRepository)

	// Controller
	membershipController := controller.NewMembershipController(membershipUsecase)
	customerController := controller.NewCustomerController(customerUsecase)
	carController := controller.NewCarController(carUsecase)
	bookingTypeController := controller.NewBookingTypeController(bookingTypeUsecase)
	driverController := controller.NewDriverController(driverUsecase)
	bookingController := controller.NewBookingController(bookingUsecase)

	routeConfig := controller.RouteConfig{
		Router:                conf.Router,
		CustomerController:    customerController,
		CarController:         carController,
		BookingTypeController: bookingTypeController,
		BookingController:     bookingController,
		MembershipController:  membershipController,
		DriverController:      driverController,
	}

	routeConfig.ConfigureRoutes()
}
