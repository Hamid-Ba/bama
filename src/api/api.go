package api

import (
	"fmt"

	"github.com/Hamid-Ba/bama/api/middlewares"
	"github.com/Hamid-Ba/bama/api/routers"
	"github.com/Hamid-Ba/bama/api/validators"
	"github.com/Hamid-Ba/bama/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()

	r := gin.New()

	val, _ := binding.Validator.Engine().(*validator.Validate)

	val.RegisterValidation("password", validators.CheckPassword, true)

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitterMiddleware())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))

}
