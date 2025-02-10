package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func setAdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/admin-list", controllers.GetAdminList)
		admin.GET("/student-list", controllers.GetStudentList)
		admin.GET("/recruiter-list", controllers.GetRecruiterList)
		admin.POST("/create-notice", controllers.CreateNotice)
		admin.GET("/notices", controllers.GetAllNotice)
		admin.POST("/create-event", controllers.CreateEvent)
		admin.PUT("/toggle-event-activation", controllers.ToggleEventActivation)
		admin.GET("/resume-list", controllers.GetResumeListForAdmin)
		admin.POST("/verify-resume", middleware.RequireAuth, controllers.VerifyResume)
		admin.DELETE("/delete-resume", middleware.RequireAuth, middleware.AdminAuth, controllers.DeleteFile)
	}
}
