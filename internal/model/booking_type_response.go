package model

import "carrental/internal/entity"

type BookingTypeResponse struct {
	ID          int    `json:"id"`
	BookingType string `json:"booking_type"`
	Description string `json:"description"`
}

func ToBookingTypeResponse(bookingType *entity.BookingType) *BookingTypeResponse {
	return &BookingTypeResponse{
		ID:          bookingType.ID,
		BookingType: bookingType.BookingType,
		Description: bookingType.Description,
	}
}

func ToBookingTypesResponse(bookingTypes []entity.BookingType) []BookingTypeResponse {
	response := make([]BookingTypeResponse, len(bookingTypes))
	for i, bookingType := range bookingTypes {
		response[i] = *ToBookingTypeResponse(&bookingType)
	}
	return response
}
