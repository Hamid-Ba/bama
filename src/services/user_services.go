package services

import (
	"github.com/Hamid-Ba/bama/common"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/infrastructure/db"
	"gorm.io/gorm"
)

type UserService struct {
	cfg         *config.Config
	otp_service *OTPSerivce
	db          *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		cfg:         cfg,
		otp_service: NewOTPService(cfg),
		db:          db.GetDb()}
}

func (user_service *UserService) SendOTP(mobile_number string) error {
	otp := common.GenerateOtp()

	err := user_service.otp_service.SetOTP(mobile_number, otp)

	if err != nil {
		return err
	}

	return nil
}
