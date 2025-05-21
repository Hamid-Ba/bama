package services

import (
	"database/sql"
	"time"

	"github.com/Hamid-Ba/bama/api/dtos"
	"github.com/Hamid-Ba/bama/domain/models"
	"github.com/Hamid-Ba/bama/infrastructure/db"
	"github.com/Hamid-Ba/bama/pkg/logging"
	"gorm.io/gorm"
)

type CountryService struct {
	db *gorm.DB
}

func NewCountryService() *CountryService {
	return &CountryService{db: db.GetDb()}
}

func (country_service *CountryService) GetBy(id int) (*dtos.CountryResponseDTO, error) {
	res := new(dtos.CountryResponseDTO)

	if err := country_service.db.Model(models.Country{}).Where("id = ? AND IsActive = ?", id, true).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (country_service *CountryService) GetList() (*[]dtos.CountryResponseDTO, error) {
	res := new([]dtos.CountryResponseDTO)

	if err := country_service.db.Model(models.Country{}).Where("IsActive = ?", true).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (country_service *CountryService) Create(create_dto dtos.CreateUpdateCountryDTO) (*dtos.CountryResponseDTO, error) {
	country := &models.Country{
		Name: create_dto.Name,
	}

	tx := country_service.db.Statement.Begin()

	country.BeforeCreate()
	if err := tx.Create(country).Error; err != nil {
		tx.Rollback()
		logging.Log.Error(err.Error())
		return nil, err
	}

	tx.Commit()

	return country_service.GetBy(country.Id)
}

func (country_service *CountryService) Update(id int, update_dto dtos.CreateUpdateCountryDTO) (*dtos.CountryResponseDTO, error) {
	updated_field := map[string]interface{}{
		"Name":       update_dto.Name,
		"Updated_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := country_service.db.Statement.Begin()
	if err := tx.Model(models.Country{}).Where("id = ?", id).Updates(updated_field).Error; err != nil {
		tx.Rollback()
		logging.Log.Error(err.Error())
		return nil, err
	}
	tx.Commit()

	return country_service.GetBy(id)
}

func (country_service *CountryService) Delete(id int) error {
	tx := country_service.db.Statement.Begin()

	if err := tx.Model(models.Country{}).Where("id = ?", id).Update("IsActive", false).Error; err != nil {
		tx.Rollback()
		logging.Log.Error(err.Error())
		return err
	}

	tx.Commit()

	return nil
}
