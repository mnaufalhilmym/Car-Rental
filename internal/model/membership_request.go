package model

type CreateMembershipRequest struct {
	MembershipName string  `json:"membership_name" binding:"required,gt=0"`
	Discount       float64 `json:"discount" binding:"required,gt=0"`
}

type DeleteMembershipRequest struct {
	ID int `uri:"id" binding:"required,gt=0"`
}
