package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (health *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working Fine!")
}
