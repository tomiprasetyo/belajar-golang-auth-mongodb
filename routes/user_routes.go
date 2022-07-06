package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-golang-auth-mongodb/controllers"
	"github.com/tomiprasetyo/belajar-golang-auth-mongodb/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users:user_id", controllers.GetUser())
}
