package model

import (
	apperror "carrental/internal/error"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response[T any] struct {
	Error      string      `json:"error,omitempty"`
	Pagination *pagination `json:"pagination,omitempty"`
	Data       T           `json:"data,omitempty"`
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

func ResponseOK[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, response[T]{
		Data: data,
	})
}

func ResponseOKPaginated[T any](ctx *gin.Context, data []T, totalItem int64, page int, size int) {
	ctx.JSON(http.StatusOK, response[[]T]{
		Data: data,
		Pagination: &pagination{
			Page:      page,
			Size:      size,
			TotalItem: totalItem,
			TotalPage: int64(math.Ceil(float64(totalItem) / float64(size))),
		},
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
