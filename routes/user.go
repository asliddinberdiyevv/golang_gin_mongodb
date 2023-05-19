package routes

import (
	"bookhub/controllers"
	"bookhub/middleware"

	"github.com/gin-gonic/gin"
)

func User(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/api/users", controllers.ListUser())
	incomingRoutes.GET("/api/users/:user_id", controllers.User())
}