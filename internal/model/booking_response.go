package model

import (
	"carrental/internal/entity"
	"time"
)

type BookingResponse struct {
	ID        int              `json:"id"`
	Customer  CustomerResponse `json:"customer"`
	Car       CarResponse      `json:"car"`
	StartRent time.Time        `json:"start_rent"`
	EndRent   time.Time        `json:"end_rent"`
	TotalCost float64          `json:"total_cost"`
	Finished  bool             `json:"finished"`
}

func ToBookingResponse(booking *entity.Booking) *BookingResponse {
	return &BookingResponse{
		ID:        booking.ID,
		Customer:  CustomerResponse(booking.Customer),
		Car:       CarResponse(booking.Car),
		StartRent: booking.StartRent,
		EndRent:   booking.EndRent,
		TotalCost: booking.TotalCost,
		Finished:  booking.Finished,
	}
}

func ToBookingsResponse(bookings []entity.Booking) []BookingResponse {
	response := make([]BookingResponse, len(bookings))
	for i, booking := range bookings {
		response[i] = *ToBookingResponse(&booking)
	}
	return response
}
