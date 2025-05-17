package routers

import (
	"github.com/Hamid-Ba/bama/api/handlers"
	"github.com/Hamid-Ba/bama/api/middlewares"
	"github.com/Hamid-Ba/bama/config"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {

	handler := handlers.NewUserHandler(cfg)

	r.POST("/send-otp", middlewares.OTPLimiter(cfg), handler.SendOTP)
	r.POST("/LoginOrRegister", middlewares.OTPLimiter(cfg), handler.LoginOrRegister)
}
