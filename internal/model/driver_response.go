package model

import "carrental/internal/entity"

type DriverResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	NIK         string  `json:"nik"`
	PhoneNumber string  `json:"phone_number"`
	DailyCost   float64 `json:"daily_cost"`
}

func ToDriverResponse(driver *entity.Driver) *DriverResponse {
	return &DriverResponse{
		ID:          driver.ID,
		Name:        driver.Name,
		NIK:         driver.NIK,
		PhoneNumber: driver.PhoneNumber,
		DailyCost:   driver.DailyCost,
	}
}

func ToDriversResponse(drivers []entity.Driver) []DriverResponse {
	response := make([]DriverResponse, len(drivers))
	for i, driver := range drivers {
		response[i] = *ToDriverResponse(&driver)
	}
	return response
}
