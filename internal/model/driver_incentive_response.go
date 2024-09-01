package model

import "carrental/internal/entity"

type DriverIncentiveResponse struct {
	ID        int     `json:"id"`
	Incentive float64 `json:"incentive"`
}

func ToDriverIncentiveResponse(driverIncentive *entity.DriverIncentive) *DriverIncentiveResponse {
	return &DriverIncentiveResponse{
		ID:        driverIncentive.ID,
		Incentive: driverIncentive.Incentive,
	}
}
