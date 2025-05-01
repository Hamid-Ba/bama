package routers

import (
	"github.com/Hamid-Ba/bama/api/handlers"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {

	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}
