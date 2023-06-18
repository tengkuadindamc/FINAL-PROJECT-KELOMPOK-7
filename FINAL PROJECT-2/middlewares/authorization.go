package middlewares

import (
	"fp-2/database"
	"fp-2/entities"
	"fp-2/pkg/errs"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()

		photoID, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		userData := c.MustGet("currentUser").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := &entities.Photo{}

		err = db.First(photo, "id = ?", photoID).Error
		if err != nil {
			missingDataErr := errs.NewNotFound("Photo not found")
			c.AbortWithStatusJSON(missingDataErr.StatusCode(), missingDataErr)
			return
		}

		if photo.UserId != userID {
			notAuthorizedErr := errs.NewNotAuthorized("You don't own this photo")
			c.AbortWithStatusJSON(notAuthorizedErr.StatusCode(), notAuthorizedErr)
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()

		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		userData := c.MustGet("currentUser").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		comment := &entities.Comment{}

		err = db.First(comment, "id = ?", commentId).Error

		if err != nil {
			notFoundError := errs.NewNotFound("Comment not found")
			c.AbortWithStatusJSON(notFoundError.StatusCode(), notFoundError)
			return
		}

		if comment.UserId != userId {
			notAuthorizedErr := errs.NewNotAuthorized("You don't own this comment")
			c.AbortWithStatusJSON(notAuthorizedErr.StatusCode(), notAuthorizedErr)
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()

		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		userData := c.MustGet("currentUser").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		socialMedia := &entities.SocialMedia{}

		err = db.First(socialMedia, "id = ?", socialMediaId).Error
		if err != nil {
			missingDataErr := errs.NewNotFound("Social media not found")
			c.AbortWithStatusJSON(missingDataErr.StatusCode(), missingDataErr)
			return
		}

		if socialMedia.UserId != userId {
			notAuthorizedErr := errs.NewNotAuthorized("You don't own this social media")
			c.AbortWithStatusJSON(notAuthorizedErr.StatusCode(), notAuthorizedErr)
			return
		}

		c.Next()
	}
}
