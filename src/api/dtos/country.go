package dtos

type CreateUpdateCountryDTO struct {
	Name string `json:"name" binding:"required,min=1,max=25"`
}

type CountryResponseDTO struct {
	Id   int
	Name string
}
