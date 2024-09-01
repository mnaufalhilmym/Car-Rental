package model

import "time"

type CreateBookingRequest struct {
	CustomerID int       `json:"customer_id" validate:"required,gt=0"`
	CarID      int       `json:"car_id" validate:"required,gt=0"`
	StartRent  time.Time `json:"start_rent" validate:"required"`
	EndRent    time.Time `json:"end_rent" validate:"required,gtfield=StartRent"`
	TotalCost  float64   `json:"total_cost" validate:"required"`
	Finished   bool      `json:"finished" validate:"required"`
}

type GetBookingRequest struct {
	ID int `validate:"required,gt=0"`
}

type GetListBookingRequest struct {
	paginationRequest

	CustomerID int
	CarID      int
	StartRent  string // e.g. gt=2006-01-02 or lte=2006-01-02
	EndRent    string // e.g. gt=2006-01-02 or lte=2006-01-02
	Timezone   string // e.g. +7 or 8:45 or -09:30
	TotalCost  string // e.g. gt=10 or gte=11 or lt=9 or lte=8
	Finished   *bool
}

type UpdateBookingRequest struct {
	ID         int        `json:"-" validate:"required,gt=0"`
	CustomerID *int       `json:"customer_id" validate:"omitempty,gt=0"`
	CarID      *int       `json:"car_id" validate:"omitempty,gt=0"`
	StartRent  *time.Time `json:"start_rent"`
	EndRent    *time.Time `json:"end_rent" validate:"omitempty,gtfield=StartRent"`
	TotalCost  *float64   `json:"total_cost"`
	Finished   *bool      `json:"finished"`
}

type DeleteBookingRequest struct {
	ID int `validate:"required,gt=0"`
}
