package model

type CreateCustomerRequest struct {
	Name         string `json:"name" binding:"required,gt=0"`
	NIK          string `json:"nik" binding:"required,numeric,gt=10"`
	PhoneNumber  string `json:"phone_number" binding:"required,phone_number"`
	MembershipID *int   `json:"membership_id" binding:"omitempty,gt=0"`
}

type GetCustomerRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}

type GetListCustomerRequest struct {
	paginationRequest

	Name        string `form:"name"`
	NIK         string `form:"nik"`
	PhoneNumber string `form:"phone_number"`
}

type UpdateCustomerRequest struct {
	ID           int     `json:"-" uri:"id" binding:"required,gt=0"`
	Name         *string `json:"name" uri:"-" binding:"omitempty,gt=0"`
	NIK          *string `json:"nik" uri:"-" binding:"omitempty,numeric,gt=10"`
	PhoneNumber  *string `json:"phone_number" uri:"-" binding:"omitempty,phone_number"`
	MembershipID *int    `json:"membership_id" uri:"-" binding:"omitempty,gt=0"`
}

type DeleteCustomerRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
