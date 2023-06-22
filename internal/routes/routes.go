package routes

import (
	"github.com/asargin-dev/soundproof-backend-demo/internal/handlers/auth"
	"github.com/asargin-dev/soundproof-backend-demo/internal/handlers/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(engine *gin.Engine) {

	authRoutes := engine.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
	}

	userRoutes := engine.Group("/user")
	{
		userRoutes.GET("/profile", user.GetUserProfile)
		userRoutes.PUT("/profile", user.UpdateUserProfile)
	}

	// Swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
