package routes

import (
	"bookhub/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/signin", controllers.Signin())
}