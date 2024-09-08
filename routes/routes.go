package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Allow requests from your Next.js frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow the necessary methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Allow credentials like cookies, etc.
		MaxAge:           12 * time.Hour, // Cache the preflight result for 12 hours
	}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	SetAuthRoutes(router)

	// set routes for remaining paths
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "this route does not exist"})
	})
	return router
}
