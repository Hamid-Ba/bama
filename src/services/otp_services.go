package services

import (
	"fmt"
	"time"

	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/infrastructure/cache"
	"github.com/redis/go-redis/v9"
)

type OTPSerivce struct {
	cfg         *config.Config
	redisClient *redis.Client
}

type OTPDTO struct {
	Password string
	IsUsed   bool
}

func NewOTPService(cfg *config.Config) *OTPSerivce {
	return &OTPSerivce{
		cfg:         cfg,
		redisClient: cache.GetRedisClient(),
	}
}

func (otp_service *OTPSerivce) SetOTP(phone string, otp string) error {
	otp_dto := &OTPDTO{
		Password: otp,
		IsUsed:   false,
	}
	key := fmt.Sprintf("%s:%s", "OTP", phone)

	res, err := cache.Get[OTPDTO](otp_service.redisClient, key)
	if err == nil && !res.IsUsed {
		return fmt.Errorf("OTP EXIST")
	} else if err == nil && res.IsUsed {
		return fmt.Errorf("OTP USED")
	}
	err = cache.Set(otp_service.redisClient, key, otp_dto, otp_service.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}
