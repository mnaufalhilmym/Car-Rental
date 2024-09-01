package controller

import (
	apperror "carrental/internal/error"
	"carrental/internal/model"
	"carrental/internal/usecase"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mnaufalhilmym/gotracing"
)

type DriverController struct {
	usecase *usecase.DriverUsecase
}

func NewDriverController(uc *usecase.DriverUsecase) *DriverController {
	return &DriverController{
		usecase: uc,
	}
}

func (c *DriverController) Create(ctx *gin.Context) {
	request := new(model.CreateDriverRequest)
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

func (c *DriverController) Get(ctx *gin.Context) {
	request := new(model.GetDriverRequest)
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

func (c *DriverController) GetList(ctx *gin.Context) {
	request := new(model.GetListDriverRequest)
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

func (c *DriverController) Update(ctx *gin.Context) {
	request := new(model.UpdateDriverRequest)
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

func (c *DriverController) Delete(ctx *gin.Context) {
	request := new(model.DeleteDriverRequest)
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
