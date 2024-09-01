package controller

import (
	apperror "carrental/internal/error"
	"carrental/internal/model"
	"carrental/internal/usecase"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mnaufalhilmym/gotracing"
)

type BookingTypeController struct {
	usecase *usecase.BookingTypeUsecase
}

func NewBookingTypeController(uc *usecase.BookingTypeUsecase) *BookingTypeController {
	return &BookingTypeController{
		usecase: uc,
	}
}

func (c *BookingTypeController) Create(ctx *gin.Context) {
	request := new(model.CreateBookingTypeRequest)
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

func (c *BookingTypeController) GetAll(ctx *gin.Context) {
	response, err := c.usecase.GetAll(ctx.Request.Context())
	if err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseOK(ctx, response)
}

func (c *BookingTypeController) Delete(ctx *gin.Context) {
	request := new(model.DeleteBookingTypeRequest)
	if err := ctx.ShouldBindUri(request); err != nil {
		gotracing.Error("Failed to parse request", err)
		model.ResponseError(ctx, apperror.BadRequest(errors.New("failed to parse request")))
		return
	}

	if err := c.usecase.Delete(ctx.Request.Context(), request); err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseOK(ctx, true)
}
