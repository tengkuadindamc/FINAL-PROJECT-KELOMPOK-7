package main

import (
	"finalproject4/config"
	"finalproject4/handler"
	"finalproject4/helper"
	"finalproject4/repository"
	"finalproject4/route"
	"finalproject4/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	if err != nil {
		panic(err)
	}

	auth := helper.NewService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, auth)

	categoryRepository := repository.NewCategoriesRepository(db)
	categoryService := service.NewCategoriesService(categoryRepository)
	categoryHandler := handler.NewCategoriesHandler(categoryService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	transactionHistoryRepository := repository.NewTransactionHistoryRepository(db)
	transactionHistoryService := service.NewTransactionHistoryService(transactionHistoryRepository)
	transactionHistoryHandler := handler.NewTransactionHistoryHandler(transactionHistoryService)

	route.TransactionRouter(router, transactionHistoryHandler, db)
	route.CategoryRouter(router, categoryHandler, db)
	route.UserRouter(router, userHandler, db)
	route.ProductRouter(router, productHandler, db)

	router.Run(":" + cfg.ServerPort)
}
