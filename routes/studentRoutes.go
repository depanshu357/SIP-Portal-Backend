package routes

import (
	"sip/controllers"
	middleware "sip/middlewares"

	"github.com/gin-gonic/gin"
)

func SetStudentRoutes(router *gin.Engine) {
	student := router.Group("/student")
	{
		student.GET("/profile-info", middleware.RequireAuth, controllers.GetStudentProfile)
		student.GET("/notices", middleware.RequireAuth, controllers.GetStudentNotice)
		student.POST("/update-profile", middleware.RequireAuth, controllers.UpdateProfile)
		student.POST("/upload-file", middleware.RequireAuth, controllers.UploadFile)
		student.GET("/info-for-resume-name", middleware.RequireAuth, controllers.GetStudentInfoForResumeName)
		student.GET("/resume-list", middleware.RequireAuth, controllers.GetResumeList)
		student.GET("/download-file", middleware.RequireAuth, controllers.DownloadFile)
	}
}
