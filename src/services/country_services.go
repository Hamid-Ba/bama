package services

import (
	"github.com/Hamid-Ba/bama/api/dtos"
	"github.com/Hamid-Ba/bama/domain/models"
	"github.com/Hamid-Ba/bama/infrastructure/db"
)

type CountryService struct {
	repo *RepositroyService[models.Country, dtos.CreateUpdateCountryDTO, dtos.CreateUpdateCountryDTO, dtos.CountryResponseDTO]
}

func NewCountryService() *CountryService {
	return &CountryService{repo: &RepositroyService[models.Country, dtos.CreateUpdateCountryDTO, dtos.CreateUpdateCountryDTO, dtos.CountryResponseDTO]{
		db: db.GetDb(),
	}}
}

func (country_service *CountryService) GetBy(id int) (*dtos.CountryResponseDTO, error) {
	res, err := country_service.repo.GetBy(id)

	return &res, err
}

func (country_service *CountryService) GetList() (*[]dtos.CountryResponseDTO, error) {
	res, err := country_service.repo.GetList()

	return res, err
}

func (country_service *CountryService) Create(create_dto dtos.CreateUpdateCountryDTO) (*dtos.CountryResponseDTO, error) {
	res, err := country_service.repo.Create(create_dto)

	return &res, err
}

func (country_service *CountryService) Update(id int, update_dto dtos.CreateUpdateCountryDTO) (*dtos.CountryResponseDTO, error) {
	res, err := country_service.repo.Update(id, update_dto)

	return &res, err
}

func (country_service *CountryService) Delete(id int) error {
	err := country_service.repo.Delete(id)
	return err
}
