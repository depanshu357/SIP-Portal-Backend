package middleware

import (
	"sip/models"
	"sip/utils"

	"github.com/gin-gonic/gin"
)

func StudentAuth(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.Logger.Sugar().Error("User not found in context")
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if userModel, ok := user.(models.User); ok {
		utils.Logger.Sugar().Infof("User: %T with email: %s", userModel, userModel.Email)

		if userModel.Role != "student" {
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
