package controller

import (
	apperror "carrental/internal/error"
	"carrental/internal/model"
	"carrental/internal/usecase"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mnaufalhilmym/gotracing"
)

type MembershipController struct {
	usecase *usecase.MembershipUsecase
}

func NewMembershipController(uc *usecase.MembershipUsecase) *MembershipController {
	return &MembershipController{
		usecase: uc,
	}
}

func (c *MembershipController) Create(ctx *gin.Context) {
	request := new(model.CreateMembershipRequest)
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

func (c *MembershipController) GetAll(ctx *gin.Context) {
	response, err := c.usecase.GetAll(ctx.Request.Context())
	if err != nil {
		model.ResponseError(ctx, err)
		return
	}

	model.ResponseOK(ctx, response)
}

func (c *MembershipController) Delete(ctx *gin.Context) {
	request := new(model.DeleteMembershipRequest)
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
