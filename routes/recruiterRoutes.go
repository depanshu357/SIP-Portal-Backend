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
		// recruiter.Get("/jobs", controllers.GetRecruiterJobs)
	}
}
