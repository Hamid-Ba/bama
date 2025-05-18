package middlewares

import (
	"net/http"

	"github.com/Hamid-Ba/bama/api/helpers"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helpers.GenerateBaseResponseWithError(nil, false, 50001, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helpers.GenerateBaseResponseWithAnyError(nil, false, 50001, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
