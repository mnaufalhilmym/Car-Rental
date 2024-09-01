package model

import "carrental/internal/entity"

type CustomerResponse struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	NIK         string              `json:"nik"`
	PhoneNumber string              `json:"phone_number"`
	Membership  *MembershipResponse `json:"membership,omitempty"`
}

func ToCustomerResponse(customer *entity.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		NIK:         customer.NIK,
		PhoneNumber: customer.PhoneNumber,
		Membership: func() *MembershipResponse {
			if customer.Membership != nil {
				return ToMembershipResponse(customer.Membership)
			}
			return nil
		}(),
	}
}

func ToCustomersResponse(customers []entity.Customer) []CustomerResponse {
	response := make([]CustomerResponse, len(customers))
	for i, customer := range customers {
		response[i] = *ToCustomerResponse(&customer)
	}
	return response
}
