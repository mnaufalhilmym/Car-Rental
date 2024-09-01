package model

import (
	apperror "carrental/internal/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response[T any] struct {
	Data       T           `json:"data,omitempty"`
	Pagination *pagination `json:"pagination,omitempty"`
	Error      string      `json:"error,omitempty"`
}

type pagination struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

func ResponseCreated[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusCreated, response[T]{
		Data: data,
	})
}

func ResponseError(ctx *gin.Context, err error) {
	appError, ok := err.(*apperror.Error)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, response[any]{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(appError.Code, response[any]{
		Error: appError.Err.Error(),
	})
}
