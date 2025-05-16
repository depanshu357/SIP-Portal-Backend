package middleware

import (
	"sip/models"

	"github.com/gin-gonic/gin"
)

func RecruiterAuth(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if userModel, ok := user.(models.User); ok {
		if userModel.Role != "recruiter" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}
	} else {
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.Next()
}
