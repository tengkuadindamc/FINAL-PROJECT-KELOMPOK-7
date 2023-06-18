package route

import (
	"finalproject4/handler"
	"finalproject4/helper"
	"finalproject4/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionRouter(router *gin.Engine, handler handler.TransactionHistoryHandler, db *gorm.DB) {
	auth := helper.NewService()
	userRepo := repository.NewUserRepository(db)

	router.POST("/transactions", auth.AuthMiddleware(auth, userRepo), handler.CreateTransactionHistory)
	router.GET("/transactions/user-transactions", auth.AuthMiddleware(auth, userRepo), handler.GetAllTransactionHistory)
	router.GET("/transactions/my-transactions", auth.AuthMiddleware(auth, userRepo), handler.GetTransactionHistoryByUserId)
	router.DELETE("/transactions/:id", auth.AuthMiddleware(auth, userRepo), handler.DeleteTransactionHistory)
}
