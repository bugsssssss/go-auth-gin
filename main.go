package main

import (
	"fmt"
	"github.com/bugsssssss/auth-gin/controllers"
	"github.com/bugsssssss/auth-gin/initizalizers"
	"github.com/bugsssssss/auth-gin/middleware"
	"github.com/bugsssssss/auth-gin/seeders"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	initizalizers.LoadEnvVariables()
	initizalizers.ConnectToDb()
	initizalizers.SyncDatabase()
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())

	products := seeders.ProductSeed(5)
	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": products,
		})
	})
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.AuthMiddleware, controllers.Validate)

	//authGroup := router.Group("/api/v1")
	//authGroup.Use(middleware.AuthMiddleware())
	//{
	//	authGroup.GET("/data", func(c *gin.Context) {
	//		c.JSON(200, gin.H{"message": "Authenticated and authorized!"})
	//	})
	//}

	// Run the server on port 8080
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
