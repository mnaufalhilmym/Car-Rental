package controller

import "github.com/gin-gonic/gin"

type RouteConfig struct {
	Router                *gin.Engine
	CustomerController    *CustomerController
	CarController         *CarController
	BookingController     *BookingController
	BookingTypeController *BookingTypeController
	MembershipController  *MembershipController
	DriverController      *DriverController
}

func (r *RouteConfig) ConfigureRoutes() {
	// V1 group
	{
		v1 := r.Router.Group("/v1")

		v1.POST("/customer", r.CustomerController.Create)
		v1.GET("/customer/:id", r.CustomerController.Get)
		v1.GET("/customers", r.CustomerController.GetList)
		v1.PATCH("/customer/:id", r.CustomerController.Update)
		v1.DELETE("/customer/:id", r.CustomerController.Delete)

		v1.POST("/car", r.CarController.Create)
		v1.GET("/car/:id", r.CarController.Get)
		v1.GET("/cars", r.CarController.GetList)
		v1.PATCH("/car/:id", r.CarController.Update)
		v1.DELETE("/car/:id", r.CarController.Delete)

		v1.POST("/booking", r.BookingController.Create)
		v1.GET("/booking/:id", r.BookingController.Get)
		v1.GET("/bookings", r.BookingController.GetList)
		v1.PATCH("/booking/:id", r.BookingController.Update)
		v1.DELETE("/booking/:id", r.BookingController.Delete)
	}

	// V2 group
	{
		v2 := r.Router.Group("/v2")

		v2.POST("/customer", r.CustomerController.Create)
		v2.GET("/customer/:id", r.CustomerController.Get)
		v2.GET("/customers", r.CustomerController.GetList)
		v2.PATCH("/customer/:id", r.CustomerController.Update)
		v2.DELETE("/customer/:id", r.CustomerController.Delete)

		v2.POST("/car", r.CarController.Create)
		v2.GET("/car/:id", r.CarController.Get)
		v2.GET("/cars", r.CarController.GetList)
		v2.PATCH("/car/:id", r.CarController.Update)
		v2.DELETE("/car/:id", r.CarController.Delete)

		v2.POST("/booking", r.BookingController.CreateV2)
		v2.GET("/booking/:id", r.BookingController.GetV2)
		v2.GET("/bookings", r.BookingController.GetList)
		v2.PATCH("/booking/:id", r.BookingController.UpdateV2)
		v2.DELETE("/booking/:id", r.BookingController.Delete)

		v2.POST("/booking-type", r.BookingTypeController.Create)
		v2.GET("/booking-types", r.BookingTypeController.GetAll)
		v2.DELETE("/booking-type/:id", r.BookingTypeController.Delete)

		v2.POST("/membership", r.MembershipController.Create)
		v2.GET("/memberships", r.MembershipController.GetAll)
		v2.DELETE("/membership/:id", r.MembershipController.Delete)

		v2.POST("/driver", r.DriverController.Create)
		v2.GET("/driver/:id", r.DriverController.Get)
		v2.GET("/drivers", r.DriverController.GetList)
		v2.PATCH("/driver/:id", r.DriverController.Update)
		v2.DELETE("/driver/:id", r.DriverController.Delete)
	}
}
