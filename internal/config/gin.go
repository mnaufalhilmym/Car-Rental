package config

import (
	apperror "carrental/internal/error"
	"carrental/internal/model"

	"github.com/gin-gonic/gin"
)

func NewGin(appMode string) *gin.Engine {
	gin.SetMode(appMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(customErrorHandler())

	return router
}

func customErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			model.ResponseError(ctx, apperror.InternalServerError(ctx.Errors[0]))
		}
	}
}
