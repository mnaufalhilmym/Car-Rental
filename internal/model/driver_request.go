package model

type CreateDriverRequest struct {
	Name        string  `json:"name" binding:"required,gt=0"`
	NIK         string  `json:"nik" binding:"required,numeric,gt=10"`
	PhoneNumber string  `json:"phone_number" binding:"required,phone_number"`
	DailyCost   float64 `json:"daily_cost" binding:"required,gt=0"`
}

type GetDriverRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}

type GetListDriverRequest struct {
	paginationRequest

	Name        string `form:"name"`
	NIK         string `form:"nik"`
	PhoneNumber string `form:"phone_number"`
	DailyCost   string `form:"daily_cost"`
}

type UpdateDriverRequest struct {
	ID          int      `json:"-" uri:"id" binding:"required,gt=0"`
	Name        *string  `json:"name" uri:"-" binding:"omitempty,gt=0"`
	NIK         *string  `json:"nik" uri:"-" binding:"omitempty,numeric,gt=10"`
	PhoneNumber *string  `json:"phone_number" uri:"-" binding:"omitempty,phone_number"`
	DailyCost   *float64 `json:"daily_cost" uri:"-" binding:"omitempty,gt=0"`
}

type DeleteDriverRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
