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
		Customer:  *ToCustomerResponse(&booking.Customer),
		Car:       *ToCarResponse(&booking.Car),
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

type BookingResponseV2 struct {
	ID              int                      `json:"id"`
	Customer        CustomerResponse         `json:"customer"`
	Car             CarResponse              `json:"car"`
	StartRent       time.Time                `json:"start_rent"`
	EndRent         time.Time                `json:"end_rent"`
	TotalCost       float64                  `json:"total_cost"`
	Finished        bool                     `json:"finished"`
	Discount        float64                  `json:"discount"`
	BookingType     *BookingTypeResponse     `json:"booking_type"`
	Driver          *DriverResponse          `json:"driver"`
	TotalDriverCost float64                  `json:"total_driver_cost"`
	DriverIncentive *DriverIncentiveResponse `json:"driver_incentive"`
}

func ToBookingResponseV2(booking *entity.Booking, driverIncentive *entity.DriverIncentive) *BookingResponseV2 {
	return &BookingResponseV2{
		ID:        booking.ID,
		Customer:  *ToCustomerResponse(&booking.Customer),
		Car:       *ToCarResponse(&booking.Car),
		StartRent: booking.StartRent,
		EndRent:   booking.EndRent,
		TotalCost: booking.TotalCost,
		Finished:  booking.Finished,
		Discount:  booking.Discount,
		BookingType: func() *BookingTypeResponse {
			if booking.BookingType != nil {
				return ToBookingTypeResponse(booking.BookingType)
			}
			return nil
		}(),
		Driver: func() *DriverResponse {
			if booking.Driver != nil {
				return ToDriverResponse(booking.Driver)
			}
			return nil
		}(),
		TotalDriverCost: booking.TotalDriverCost,
		DriverIncentive: func() *DriverIncentiveResponse {
			if driverIncentive != nil {
				return ToDriverIncentiveResponse(driverIncentive)
			}
			return nil
		}(),
	}
}

func ToBookingsResponseV2(bookings []entity.Booking) []BookingResponseV2 {
	response := make([]BookingResponseV2, len(bookings))
	for i, booking := range bookings {
		response[i] = *ToBookingResponseV2(&booking, nil)
	}
	return response
}
