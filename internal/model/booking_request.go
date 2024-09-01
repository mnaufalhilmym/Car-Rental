package model

import "time"

type CreateBookingRequest struct {
	CustomerID int       `json:"customer_id" binding:"required,gt=0"`
	CarID      int       `json:"car_id" binding:"required,gt=0"`
	StartRent  time.Time `json:"start_rent" binding:"required"`
	EndRent    time.Time `json:"end_rent" binding:"required,gtfield=StartRent"`
	TotalCost  float64   `json:"total_cost" binding:"required"`
	Finished   bool      `json:"finished" binding:"required"`
}

type GetBookingRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}

type GetListBookingRequest struct {
	paginationRequest

	CustomerID int    `form:"customer_id" binding:"omitempty,gt=0"`
	CarID      int    `form:"car_id" binding:"omitempty,gt=0"`
	StartRent  string `form:"start_rent"` // e.g. gt=2006-01-02 or lte=2006-01-02
	EndRent    string `form:"end_rent"`   // e.g. gt=2006-01-02 or lte=2006-01-02
	Timezone   string `form:"timezone"`   // e.g. +7 or 8:45 or -09:30
	TotalCost  string `form:"total_cost"` // e.g. gt=10 or gte=11 or lt=9 or lte=8
	Finished   *bool  `form:"finished"`
}

type UpdateBookingRequest struct {
	ID         int        `json:"-" uri:"id" binding:"required,gt=0"`
	CustomerID *int       `json:"customer_id" uri:"-" binding:"omitempty,gt=0"`
	CarID      *int       `json:"car_id" uri:"-" binding:"omitempty,gt=0"`
	StartRent  *time.Time `json:"start_rent" uri:"-"`
	EndRent    *time.Time `json:"end_rent" uri:"-" binding:"omitempty,gtfield=StartRent"`
	TotalCost  *float64   `json:"total_cost" uri:"-"`
	Finished   *bool      `json:"finished" uri:"-"`
}

type DeleteBookingRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
