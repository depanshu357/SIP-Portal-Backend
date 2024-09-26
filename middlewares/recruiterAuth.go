package middleware

import (
	"sip/models"
	"sip/utils"

	"github.com/gin-gonic/gin"
)

func RecruiterAuth(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.Logger.Sugar().Error("User not found in context")
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if userModel, ok := user.(models.User); ok {
		if userModel.Role != "recruiter" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}
	} else {
		utils.Logger.Sugar().Error("Failed to cast user to models.User")
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.Next()
}
