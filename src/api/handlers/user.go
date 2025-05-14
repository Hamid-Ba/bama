package handlers

import (
	"net/http"

	"github.com/Hamid-Ba/bama/api/dtos"
	"github.com/Hamid-Ba/bama/api/helpers"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user_service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	user_service := services.NewUserService(cfg)
	return &UserHandler{
		user_service: user_service,
	}
}

func (h *UserHandler) SendOTP(c *gin.Context) {
	dto := new(dtos.UserOTPDTO)
	err := c.ShouldBindJSON(dto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	err = h.user_service.SendOTP(dto.MobileNumber)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithError(nil, false, -1, err))
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(dto, true, 0))
}
