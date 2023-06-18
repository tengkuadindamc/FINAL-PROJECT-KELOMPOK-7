package route

import (
	"finalproject4/handler"
	"finalproject4/helper"
	"finalproject4/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRouter(router *gin.Engine, handler handler.CategoriesHandler, db *gorm.DB) {
	auth := helper.NewService()
	userRepo := repository.NewUserRepository(db)

	router.POST("/categories", auth.AuthMiddleware(auth, userRepo), handler.CreateCategory)
	router.GET("/categories", auth.AuthMiddleware(auth, userRepo), handler.GetAllCategory)
	router.PATCH("/categories/:id", auth.AuthMiddleware(auth, userRepo), handler.PatchCategory)
	router.DELETE("/categories/:id", auth.AuthMiddleware(auth, userRepo), handler.DeleteCategory)
}
