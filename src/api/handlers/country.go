package handlers

import (
	"net/http"
	"strconv"

	"github.com/Hamid-Ba/bama/api/dtos"
	"github.com/Hamid-Ba/bama/api/helpers"
	"github.com/Hamid-Ba/bama/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	country_service services.CountryService
}

func NewCountryHandler() *CountryHandler {
	return &CountryHandler{country_service: *services.NewCountryService()}
}

func (handler *CountryHandler) GetBy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	res, err := handler.country_service.GetBy(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler *CountryHandler) List(ctx *gin.Context) {
	res, err := handler.country_service.GetList()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler *CountryHandler) Create(ctx *gin.Context) {
	res_dto := new(dtos.CreateUpdateCountryDTO)

	err := ctx.ShouldBindJSON(&res_dto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	res, err := handler.country_service.Create(*res_dto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (handler *CountryHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	res_dto := new(dtos.CreateUpdateCountryDTO)
	err = ctx.ShouldBindJSON(&res_dto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	res, err := handler.country_service.Update(id, *res_dto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler *CountryHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	err = handler.country_service.Delete(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
