package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.Engine) {
	auth := router.Group("/")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
		auth.GET("/validate", middleware.RequireAuth, controllers.Validate)
		auth.POST("/send-otp", controllers.GenerateAndSendOTP)
		auth.POST("/verify-otp", controllers.VerifyOTP)
	}

}
