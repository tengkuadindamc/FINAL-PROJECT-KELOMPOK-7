package route

import (
	"finalproject4/handler"
	"finalproject4/helper"
	"finalproject4/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouter(router *gin.Engine, handler handler.ProductHandler, db *gorm.DB) {
	auth := helper.NewService()
	userRepo := repository.NewUserRepository(db)

	router.POST("/products", auth.AuthMiddleware(auth, userRepo), handler.AddProduct)
	router.GET("/products", auth.AuthMiddleware(auth, userRepo), handler.ViewProduct)
	router.PUT("/products/:id", auth.AuthMiddleware(auth, userRepo), handler.EditProduct)
	router.DELETE("/products/:id", auth.AuthMiddleware(auth, userRepo), handler.DeleteProduct)
}
