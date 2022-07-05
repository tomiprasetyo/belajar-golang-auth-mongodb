package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-golang-auth-mongodb/routes"
	"net/http"
	"os"
)

func main() {
	// environment variable
	port = os.Getenv("PORT")

	if port == "" {
		port = "8888"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": "Access granted for api-1",
		})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": "Access granted for api-2",
		})
	})

	router.Run(":" + port)
}
