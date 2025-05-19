package routers

import (
	"github.com/Hamid-Ba/bama/api/handlers"
	"github.com/gin-gonic/gin"
)

func CountryRouter(router *gin.RouterGroup) {
	handler := handlers.NewCountryHandler()

	router.GET("/", handler.List)
	router.GET("/:id", handler.GetBy)
	router.POST("/", handler.Create)
	router.PUT("/:id", handler.Update)
	router.DELETE("/:id", handler.Delete)
}
