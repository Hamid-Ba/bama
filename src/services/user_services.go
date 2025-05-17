package services

import (
	"github.com/Hamid-Ba/bama/api/dtos"
	"github.com/Hamid-Ba/bama/common"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/domain/models"
	"github.com/Hamid-Ba/bama/infrastructure/db"
	"gorm.io/gorm"
)

type UserService struct {
	cfg           *config.Config
	otp_service   *OTPService
	token_service *TokenService
	db            *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		cfg:           cfg,
		otp_service:   NewOTPService(cfg),
		token_service: NewTokenService(cfg),
		db:            db.GetDb()}
}

func (user_service *UserService) SendOTP(mobile_number string) error {
	otp := common.GenerateOtp()

	err := user_service.otp_service.SetOTP(mobile_number, otp)

	if err != nil {
		return err
	}

	return nil
}

func (user_service *UserService) LoginOrRegister(mobile_number, otp string) (*dtos.TokenDetailDTO, error) {
	err := user_service.otp_service.ValidateOTP(mobile_number, otp)

	if err != nil {
		return nil, err
	}

	is_exist, err := user_service.existsByMobileNumber(mobile_number)

	if err != nil {
		return nil, err
	}

	u := new(models.User)
	if is_exist {
		// fetch
		user, err := user_service.fetchUserBy(mobile_number)

		if err != nil {
			return nil, err
		}

		u = user
	} else {
		// create
		user, err := user_service.createUserBy(mobile_number)

		if err != nil {
			return nil, err
		}

		u = user
	}

	token_dto := tokenDTO{UserId: u.Id, Phone: mobile_number, Fullname: u.Fullname}

	if len(u.UserRoles) > 0 {
		for _, r := range u.UserRoles {
			token_dto.Roles = append(token_dto.Roles, r.Role.Name)
		}
	}

	return user_service.token_service.GenerateToken(token_dto)

}

func (user_service *UserService) existsByMobileNumber(mobile_number string) (exists bool, err error) {
	err = user_service.db.Model(&models.User{}).Select("count(*) > 0").Where("phone = ?", mobile_number).Scan(&exists).Error

	if err != nil {
		return false, err
	}

	return
}

func (user_service *UserService) fetchUserBy(mobile_number string) (user *models.User, err error) {
	err = user_service.db.Model(&models.User{}).Preload("UserRoles.Role").Where("phone = ?", mobile_number).First(&user).Error

	if err != nil {
		return nil, err
	}

	return
}

func (user_service *UserService) createUserBy(mobile_number string) (user *models.User, err error) {
	user = &models.User{
		Fullname: mobile_number,
		Phone:    mobile_number,
	}

	tx := user_service.db.Begin()

	err = tx.Create(&user).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	default_role, err := user_service.getDefaultRole()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Create(&models.UserRole{RoleId: default_role, UserId: user.Id}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	user, err = user_service.fetchUserBy(mobile_number)

	if err != nil {
		return nil, err
	}

	return
}

func (user_service *UserService) getDefaultRole() (roleId int, err error) {

	if err = user_service.db.Model(&models.Role{}).
		Select("id").
		Where("name = ?", "Default").
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
