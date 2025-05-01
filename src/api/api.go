package api

import (
	"fmt"

	"github.com/Hamid-Ba/bama/api/routers"
	"github.com/Hamid-Ba/bama/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))

}
