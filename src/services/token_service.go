package services

import (
	"fmt"
	"time"

	"github.com/Hamid-Ba/bama/api/dtos"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/constants"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	cfg *config.Config
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

type tokenDTO struct {
	UserId   int
	Phone    string
	Fullname string
	Roles    []string
}

func (token_service *TokenService) GenerateToken(token_dto tokenDTO) (*dtos.TokenDetailDTO, error) {
	td := &dtos.TokenDetailDTO{}
	td.AccessTokenExpireTime = time.Now().Add(token_service.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(token_service.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[constants.UserIdKey] = token_dto.UserId
	atc[constants.FullnameKey] = token_dto.Fullname
	atc[constants.MobileNumberKey] = token_dto.Phone
	atc[constants.RolesKey] = token_dto.Roles
	atc[constants.ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(token_service.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}

	rtc[constants.UserIdKey] = token_dto.UserId
	rtc[constants.ExpireTimeKey] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(token_service.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (token_service *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("UNEXPECTED ERROR")
		}
		return []byte(token_service.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (token_service *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := token_service.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, fmt.Errorf("CLAIMS NOT FOUND")
}
