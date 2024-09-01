package model

type CreateCarRequest struct {
	Name      string  `json:"name" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	DailyRent float64 `json:"daily_rent" binding:"required"`
}

type GetCarRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}

type GetListCarRequest struct {
	paginationRequest

	Name      string `form:"name"`
	Stock     string `form:"stock"`      // e.g. gt=10 or gte=11 or lt=9 or lte=8
	DailyRent string `form:"daily_rent"` // e.g. gt=10 or gte=11 or lt=9 or lte=8
}

type UpdateCarRequest struct {
	ID        int      `json:"-" uri:"id" binding:"required,gt=0"`
	Name      *string  `json:"name" uri:"-" binding:"omitempty,gt=0"`
	Stock     *int     `json:"stock" uri:"-"`
	DailyRent *float64 `json:"daily_rent" uri:"-"`
}

type DeleteCarRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
