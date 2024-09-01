package controller

import (
	apperror "carrental/internal/error"
	"carrental/internal/model"
	"carrental/internal/usecase"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mnaufalhilmym/gotracing"
)

type BookingController struct {
	usecase *usecase.BookingUsecase
}

func NewBookingController(uc *usecase.BookingUsecase) *BookingController {
	return &BookingController{
		usecase: uc,
	}
}

func (c *BookingController) Create(ctx *gin.Context) {
	request := new(model.CreateBookingRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		gotracing.Error("Failed to parse request", err)
		model.ResponseError(ctx, apperror.BadRequest(errors.New("failed to parse request")))
		return
	}

	response, err := c.usecase.Create(ctx.Request.Context(), request)
	if err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseCreated(ctx, response)
}
