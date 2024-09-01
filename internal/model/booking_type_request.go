package model

type CreateBookingTypeRequest struct {
	BookingType string `json:"booking_type" binding:"required,gt=0"`
	Description string `json:"description" binding:"required,gt=0"`
}

type DeleteBookingTypeRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
