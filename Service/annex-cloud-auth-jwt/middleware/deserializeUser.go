package middleware

import (
	"annex-cloud-auth-jwt/config"
	"annex-cloud-auth-jwt/repository"
	"annex-cloud-auth-jwt/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.IUserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
        claims, err := utils.ValidateToken(token, config.TokenSecret)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
            return
        }

        // Set user role in the context
        ctx.Set("currentRole", claims.Role)

        ctx.Next()
	}
}
