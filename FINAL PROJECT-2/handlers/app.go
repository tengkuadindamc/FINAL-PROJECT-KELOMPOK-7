package handlers

import (
	"fp-2/database"
	"fp-2/repositories"

	// _ "fp-2/docs"
	"fp-2/handlers/http_handlers"
	"fp-2/middlewares"
	"fp-2/services"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
)

const port = ":8080"

func StartApp() {
	database.StartDB()
	db := database.GetPostgresInstance()

	router := gin.Default()

	userRepo := repositories.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := http_handlers.NewUserHandler(userService)

	photoRepo := repositories.NewPhotoPG(db)
	photoService := services.NewPhotoService(photoRepo, userRepo)
	photoHandler := http_handlers.NewPhotoHandler(photoService)

	commentRepo := repositories.NewCommentPG(db)
	commentService := services.NewCommentService(commentRepo, photoRepo, userRepo)
	commentHandler := http_handlers.NewCommentHandler(commentService)

	socialMediaRepo := repositories.NewSocialMediaPG(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo, userRepo)
	socialMediaHandler := http_handlers.NewSocialMediaHandler(socialMediaService)

	usersRouter := router.Group("/users")
	{
		usersRouter.POST("/register", userHandler.RegisterUser)
		usersRouter.POST("/login", userHandler.LoginUser)
		usersRouter.PUT("/:id", middlewares.Authentication(), userHandler.UpdateUser)
		usersRouter.DELETE("/", middlewares.Authentication(), userHandler.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	photoRouter.Use(middlewares.Authentication())
	{
		photoRouter.POST("/", photoHandler.CreatePhoto)
		photoRouter.GET("/", photoHandler.GetAllPhotos)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), photoHandler.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	commentRouter.Use(middlewares.Authentication())
	{
		commentRouter.POST("/", commentHandler.CreateComment)
		commentRouter.GET("/", commentHandler.GetAllComment)
		commentRouter.GET("/user/:userId", commentHandler.GetCommentsByUserId)
		commentRouter.GET("/photo/:photoId", commentHandler.GetCommentsByPhotoId)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), commentHandler.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), commentHandler.DeleteComment)
	}

	socialMediasRouter := router.Group("/socialmedias")
	socialMediasRouter.Use(middlewares.Authentication())
	{
		socialMediasRouter.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMediasRouter.GET("/", socialMediaHandler.GetAllSocialMedias)
		socialMediasRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), socialMediaHandler.UpdateSocialMedia)
		socialMediasRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMedia)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(port)
}
