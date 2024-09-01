package model

type CreateCarRequest struct {
	Name      string  `json:"name" validate:"required"`
	Stock     int     `json:"stock" validate:"required"`
	DailyRent float64 `json:"daily_rent" validate:"required"`
}

type GetCarRequest struct {
	ID int `validate:"required,gt=0"`
}

type GetListCarRequest struct {
	paginationRequest

	Name      string
	Stock     string // e.g. gt=10 or gte=11 or lt=9 or lte=8
	DailyRent string // e.g. gt=10 or gte=11 or lt=9 or lte=8
}

type UpdateCarRequest struct {
	ID        int      `json:"-" validate:"required,gt=0"`
	Name      *string  `json:"name" validate:"omitempty,gt=0"`
	Stock     *int     `json:"stock"`
	DailyRent *float64 `json:"daily_rent"`
}

type DeleteCarRequest struct {
	ID int `validate:"required,gt=0"`
}
