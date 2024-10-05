package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetRecruiterRoutes(router *gin.Engine) {
	recruiter := router.Group("/recruiter")
	{
		recruiter.GET("/profile", middleware.RequireAuth, middleware.RecruiterAuth, controllers.GetRecruiterProfile)
		recruiter.GET("/notices", middleware.RequireAuth, middleware.RecruiterAuth, controllers.GetRecruiterNotice)
		// recruiter.Get("/jobs", controllers.GetRecruiterJobs)
	}
}
