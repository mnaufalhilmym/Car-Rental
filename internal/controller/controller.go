package controller

import "github.com/gin-gonic/gin"

type RouteConfig struct {
	Router             *gin.Engine
	CustomerController *CustomerController
	CarController      *CarController
	BookingController  *BookingController
}

func (r *RouteConfig) ConfigureRoutes() {
	v1 := r.Router.Group("/v1")
	v1.POST("/customer", r.CustomerController.Create)
	v1.POST("/car", r.CarController.Create)
}
