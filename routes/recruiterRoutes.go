package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRecruiterRoutes(router *gin.Engine) {
	recruiter := router.Group("/recruiter")
	{
		recruiter.GET("/profile", middleware.RequireAuth, controllers.GetRecruiterProfile)
		recruiter.GET("/notices", middleware.RequireAuth, controllers.GetRecruiterNotice)
		recruiter.POST("/update-profile", middleware.RequireAuth, controllers.UpdateRecruiterProfile)
		recruiter.POST("/create-job", middleware.RequireAuth, controllers.CreateJobDescription)
		recruiter.GET("/job-descriptions", middleware.RequireAuth, controllers.GetJobDescriptions)
		recruiter.POST("/edit-job-description", middleware.RequireAuth, controllers.EditJobDescription)
		recruiter.GET("/get-applicants", middleware.RequireAuth, controllers.GetListOfAppliedCandidates)
	}
}
