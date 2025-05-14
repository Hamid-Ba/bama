package dtos

type UserOTPDTO struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}
