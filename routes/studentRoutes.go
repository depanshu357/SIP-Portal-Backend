package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetStudentRoutes(router *gin.Engine) {
	student := router.Group("/student")
	{
		student.GET("/profile", middleware.RequireAuth, middleware.StudentAuth, controllers.GetStudentProfile)
		// 	// student.Get("/resume", controllers.GetStudentResume)
	}
}
