package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetStudentRoutes(router *gin.Engine) {
	student := router.Group("/student")
	{
		student.GET("/profile", middleware.RequireAuth, controllers.GetStudentProfile)
		student.GET("/notices", middleware.RequireAuth, controllers.GetStudentNotice)
		student.POST("/update-profile", middleware.RequireAuth, controllers.UpdateProfile)
		// 	// student.Get("/resume", controllers.GetStudentResume)
	}
}
