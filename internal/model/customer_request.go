package model

type CreateCustomerRequest struct {
	Name        string `json:"name" validate:"required,gt=0"`
	NIK         string `json:"nik" validate:"required,numeric,gt=10"`
	PhoneNumber string `json:"phone_number" validate:"required,phone_number"`
}

type GetCustomerRequest struct {
	ID int `validate:"required,gt=0"`
}

type GetListCustomerRequest struct {
	paginationRequest

	Name        string
	NIK         string
	PhoneNumber string
}

type UpdateCustomerRequest struct {
	ID          int     `json:"-" validate:"required,gt=0"`
	Name        *string `json:"name" validate:"omitempty,gt=0"`
	NIK         *string `json:"nik" validate:"numeric,gt=10"`
	PhoneNumber *string `json:"phone_number" validate:"phone_number"`
}

type DeleteCustomerRequest struct {
	ID int `validate:"required,gt=0"`
}
