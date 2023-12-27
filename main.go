package main

import (
	"github.com/bugsssssss/auth-gin/middleware"
	"github.com/bugsssssss/auth-gin/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())

	authGroup := router.Group("/api/v1")
	authGroup.Use(middleware.AuthMiddleware())

	{
		authGroup.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
	}

	products := seeders.ProductSeed(10000)

	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": products,
		})
	})

	// Run the server on port 8080
	router.Run(":8000")
}
