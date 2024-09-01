package model

import "carrental/internal/entity"

type CarResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	DailyRent float64 `json:"daily_rent"`
}

func ToCarResponse(car *entity.Car) *CarResponse {
	return &CarResponse{
		ID:        car.ID,
		Name:      car.Name,
		Stock:     car.Stock,
		DailyRent: car.DailyRent,
	}
}

func ToCarsResponse(cars []entity.Car) []CarResponse {
	response := make([]CarResponse, len(cars))
	for i, car := range cars {
		response[i] = *ToCarResponse(&car)
	}
	return response
}
