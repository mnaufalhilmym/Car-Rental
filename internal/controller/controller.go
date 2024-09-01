package controller

import "github.com/gin-gonic/gin"

type RouteConfig struct {
	Router             *gin.Engine
	CustomerController *CustomerController
	CarController      *CarController
	BookingController  *BookingController
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
}
