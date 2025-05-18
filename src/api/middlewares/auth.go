package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Hamid-Ba/bama/api/helpers"
	"github.com/Hamid-Ba/bama/config"
	"github.com/Hamid-Ba/bama/constants"
	"github.com/Hamid-Ba/bama/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthenticationMiddleware(cfg *config.Config) gin.HandlerFunc {
	token_service := services.NewTokenService(cfg)
	return func(ctx *gin.Context) {
		auth_token := strings.Split(ctx.Request.Header.Get(constants.AuthenticationHeader), " ")

		if len(auth_token) == 0 || len(auth_token) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.GenerateBaseResponseWithValidationError(nil, false, -2, fmt.Errorf("Unauthorized")))
			return
		}

		_, err := token_service.VerifyToken(auth_token[1])

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.GenerateBaseResponseWithError(nil, false, -2, err))
			return
		}

		claimsMap := map[string]interface{}{}
		claimsMap, err = token_service.GetClaims(auth_token[1])

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				err = fmt.Errorf("TOKEN HAS BEEN EXPIRED")
			default:
				err = fmt.Errorf("INVALID TOKEN")
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.GenerateBaseResponseWithError(nil, false, -2, err))
			return
		}

		ctx.Set(constants.UserIdKey, claimsMap[constants.UserIdKey])
		ctx.Set(constants.MobileNumberKey, claimsMap[constants.MobileNumberKey])
		ctx.Set(constants.RolesKey, claimsMap[constants.RolesKey])
		ctx.Set(constants.ExpireTimeKey, claimsMap[constants.ExpireTimeKey])

		ctx.Next()
	}
}

func AuthorizationMiddleware(role_names []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseResponse(nil, false, -2))
			return
		}
		rolesVal := ctx.Keys[constants.RolesKey]
		if rolesVal == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseResponse(nil, false, -2))
			return
		}

		roles := rolesVal.([]interface{})

		val := map[string]int{}

		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, role := range role_names {
			if _, ok := val[role]; ok {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseResponse(nil, false, -2))
	}
}
