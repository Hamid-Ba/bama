package middlewares

import (
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/Hamid-Ba/bama/api/helpers"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OTPLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(getIP(c.Request.RemoteAddr))
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helpers.GenerateBaseResponseWithError(nil, false, -1, errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func getIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return ip
}
