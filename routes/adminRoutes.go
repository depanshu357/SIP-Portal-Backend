package routes

import (
	"sip/controllers"

	"github.com/gin-gonic/gin"
)

func setAdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/admin-list", controllers.GetAdminList)
		admin.GET("/student-list", controllers.GetStudentList)
		admin.GET("/recruiter-list", controllers.GetRecruiterList)
	}
}
