package model

import "carrental/internal/entity"

type MembershipResponse struct {
	ID             int     `json:"id"`
	MembershipName string  `json:"membership_name"`
	Discount       float64 `json:"discount"`
}

func ToMembershipResponse(membership *entity.Membership) *MembershipResponse {
	return &MembershipResponse{
		ID:             membership.ID,
		MembershipName: membership.MembershipName,
		Discount:       membership.Discount,
	}
}

func ToMembershipsResponse(memberships []entity.Membership) []MembershipResponse {
	response := make([]MembershipResponse, len(memberships))
	for i, membership := range memberships {
		response[i] = *ToMembershipResponse(&membership)
	}
	return response
}
