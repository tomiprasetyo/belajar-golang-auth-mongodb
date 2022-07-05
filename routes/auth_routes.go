package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-golang-auth-mongodb/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/singup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
}
