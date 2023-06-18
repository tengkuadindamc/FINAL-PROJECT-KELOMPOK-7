package route

import (
	"finalproject4/handler"
	"finalproject4/helper"
	"finalproject4/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, handler handler.UserHandler, db *gorm.DB) {
	auth := helper.NewService()
	userRepo := repository.NewUserRepository(db)

	router.POST("user/register", handler.RegisterUser)
	router.POST("user/login", handler.LoginUser)
	router.PATCH("/user/topup", auth.AuthMiddleware(auth, userRepo), handler.UpdateBalance)

}
