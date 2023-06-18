package auth

import (
	jwt_local "final-project3/utils/jwt"

	"github.com/gin-gonic/gin"
)

func MiddlewareLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate jwt token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(401)
			return
		}

		claims, err := jwt_local.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_info", claims)
		c.Next()
	}
}
