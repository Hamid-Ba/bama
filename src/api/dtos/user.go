package dtos

type UserOTPDTO struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type TokenDetailDTO struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpireTime  int64
	RefreshTokenExpireTime int64
}

type LoginRegisterDTO struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	OTP   string `json:"otp" binding:"required,min=6,max=6"`
}
