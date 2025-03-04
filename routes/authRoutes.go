package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.Engine) {
	auth := router.Group("/")
	{
		auth.POST("/sign-up", controllers.Signup)
		auth.POST("/log-in", controllers.Login)
		auth.GET("/validate", middleware.RequireAuth, controllers.Validate)
		auth.POST("/send-otp", controllers.GenerateAndSendOTP)
		auth.POST("/verify-otp", controllers.VerifyOTP)
		auth.PUT("/change-password", controllers.ChangePassword)
		auth.GET("/download-file", middleware.RequireAuth, controllers.DownloadFile)
	}
}
