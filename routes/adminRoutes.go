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
		admin.GET("/student-list-for-event", controllers.GetStudentListForEvent)
		admin.GET("/recruiter-list", controllers.GetRecruiterList)
		admin.POST("/create-notice", controllers.CreateNotice)
		admin.DELETE("/delete-notice", middleware.RequireAuth, middleware.AdminAuth, controllers.DeleteNotice)
		admin.GET("/notices", controllers.GetAllNotice)
		admin.POST("/create-event", controllers.CreateEvent)
		admin.PUT("/toggle-event-activation", controllers.ToggleEventActivation)
		admin.GET("/resume-list", controllers.GetResumeListForAdmin)
		admin.POST("/verify-resume", middleware.RequireAuth, controllers.VerifyResume)
		admin.DELETE("/delete-resume", middleware.RequireAuth, controllers.DeleteFile)
		admin.GET("/proforma", middleware.RequireAuth, controllers.GetProforma)
		admin.GET("/job-descriptions", middleware.RequireAuth, controllers.GetAllJobDescriptions)
		admin.PUT("/toggle-proforma-visibility", middleware.RequireAuth, middleware.AdminAuth, controllers.ToggleProformaVisibility)
		admin.GET("/get-applicants", middleware.RequireAuth, controllers.GetListOfAppliedCandidates)
		admin.POST("/change-admin-access", middleware.RequireAuth, middleware.AdminAuth, controllers.ChangeAdminAccess)
		admin.POST("/change-profile-verification", middleware.RequireAuth, middleware.AdminAuth, controllers.ChangeProfileVerification)
		admin.POST("/toggle-verification-for-event", middleware.RequireAuth, controllers.ToggleVerificationForEvent)
		admin.POST("/toggle-freezing-for-event", middleware.RequireAuth, controllers.ToggleFreezingForEvent)
	}
}
