package routes

import (
	"bookhub/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("auth/signup", controllers.Signup())
	incomingRoutes.POST("auth/signin", controllers.Signin())
}