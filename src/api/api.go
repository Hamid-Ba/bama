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

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidator()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitterMiddleware())

	RegisterRouter(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))

}

func RegisterRouter(r *gin.Engine, cfg *config.Config) {
	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)

		user := v1.Group("/user")
		routers.UserRouter(user, cfg)
	}
}

func RegisterValidator() {
	val, _ := binding.Validator.Engine().(*validator.Validate)

	// Register custom validators here
	val.RegisterValidation("password", validators.CheckPassword, true)
}
