package routes

import (
	"bookhub/controllers"
	"bookhub/middleware"

	"github.com/gin-gonic/gin"
)

func User(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/users", controllers.ListUser())
	incomingRoutes.GET("/users/:user_id", controllers.User())
}