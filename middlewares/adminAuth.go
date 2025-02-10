package middleware

import (
	"github.com/gin-gonic/gin"
)

func AdminAuth(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}
