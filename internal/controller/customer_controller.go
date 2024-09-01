package controller

import (
	apperror "carrental/internal/error"
	"carrental/internal/model"
	"carrental/internal/usecase"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mnaufalhilmym/gotracing"
)

type CustomerController struct {
	usecase *usecase.CustomerUsecase
}

func NewCustomerController(uc *usecase.CustomerUsecase) *CustomerController {
	return &CustomerController{
		usecase: uc,
	}
}

func (c *CustomerController) Create(ctx *gin.Context) {
	request := new(model.CreateCustomerRequest)
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

func (c *CustomerController) Get(ctx *gin.Context) {
	request := new(model.GetCustomerRequest)
	if err := ctx.ShouldBindUri(request); err != nil {
		gotracing.Error("Failed to parse request", err)
		model.ResponseError(ctx, apperror.BadRequest(errors.New("failed to parse request")))
		return
	}

	response, err := c.usecase.Get(ctx.Request.Context(), request)
	if err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseOK(ctx, response)
}

func (c *CustomerController) GetList(ctx *gin.Context) {
	request := new(model.GetListCustomerRequest)
	if err := ctx.ShouldBindQuery(request); err != nil {
		gotracing.Error("Failed to parse request", err)
		model.ResponseError(ctx, apperror.BadRequest(errors.New("failed to parse request")))
		return
	}

	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Size <= 0 {
		request.Size = 10
	}

	response, total, err := c.usecase.GetList(ctx.Request.Context(), request)
	if err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseOKPaginated(ctx, response, total, request.Page, request.Size)
}

func (c *CustomerController) Update(ctx *gin.Context) {
	request := new(model.UpdateCustomerRequest)
	if err := ctx.ShouldBindUri(request); err != nil {
		gotracing.Error("Failed to parse request", err)
		model.ResponseError(ctx, apperror.BadRequest(errors.New("failed to parse request")))
		return
	}
	if err := ctx.ShouldBindJSON(request); err != nil {
		gotracing.Error("Failed to parse request", err)
		model.ResponseError(ctx, apperror.BadRequest(errors.New("failed to parse request")))
		return
	}

	response, err := c.usecase.Update(ctx.Request.Context(), request)
	if err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseOK(ctx, response)
}

func (c *CustomerController) Delete(ctx *gin.Context) {
	request := new(model.DeleteCustomerRequest)
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
