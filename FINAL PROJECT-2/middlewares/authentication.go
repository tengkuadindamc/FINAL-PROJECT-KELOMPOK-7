package middlewares

import (
	"fp-2/database"
	"fp-2/entities"
	"fp-2/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		data := verifyToken.(jwt.MapClaims)

		id := uint(data["id"].(float64))
		if _, isExist := data["username"]; !isExist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "Please log in again",
			})
			return
		}
		username := data["username"].(string)

		db := database.GetPostgresInstance()
		userFromDB := &entities.User{}
		err = db.First(userFromDB, "id = ?", id).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "Please log in again",
			})
			return
		}

		if userFromDB.Username != username {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "Please log in again",
			})
			return
		}

		c.Set("currentUser", verifyToken)
		c.Next()
	}
}
