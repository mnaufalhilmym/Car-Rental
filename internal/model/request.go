package model

type paginationRequest struct {
	Page int `validate:"omitempty,gt=0"`
	Size int `validate:"omitempty,gt=0"`
}
