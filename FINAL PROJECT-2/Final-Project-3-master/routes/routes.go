package routes

import (
	"final-project3/pkg/category"
	"final-project3/pkg/task"
	"final-project3/pkg/user"
	"final-project3/utils/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHttpRoute(g *gin.Engine, db *gorm.DB) {

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userController := user.InitHttpUserController(db)
	userGroup := g.Group("users")
	userGroup.POST("/register", userController.Register)
	userGroup.POST("/login", userController.Login)
	userGroup.Use(auth.MiddlewareLogin())
	{
		userGroup.PUT("/update-account/:id", userController.UpdateUserById)
		userGroup.DELETE("/delete-account/:id", userController.DeleteUserById)
	}

	categoryController := category.InitHttpCategoryController(db)
	categoryGroup := g.Group("categories")
	categoryGroup.Use(auth.MiddlewareLogin())
	{
		categoryGroup.POST("/", categoryController.CreateNewCategory)
		categoryGroup.GET("/", categoryController.GetAllCategory)
		categoryGroup.PATCH("/:categoryId", categoryController.UpdateCategoryById)
		categoryGroup.DELETE("/:categoryId", categoryController.DeleteCategoryById)
	}

	taskController := task.InitHttpTaskController(db)
	taskGroup := g.Group("tasks")
	taskGroup.Use(auth.MiddlewareLogin())
	{
		taskGroup.POST("/", taskController.CreateNewTask)
		taskGroup.GET("/", taskController.GetAllTask)
		taskGroup.PUT("/:taskId", taskController.UpdateTaskById)
		taskGroup.PATCH("/update-status/:taskId", taskController.UpdateStatusByTaskId)
		taskGroup.PATCH("/update-category/:taskId", taskController.UpdateCategoryByTaskId)
		taskGroup.DELETE("/:taskId", taskController.DeleteTaskById)

	}

}
