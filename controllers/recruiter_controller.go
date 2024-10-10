package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"

	"github.com/gin-gonic/gin"
)

func GetRecruiterProfile(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var existingUser models.Recruiter
	if err := database.DB.Where("user_id = ?", user_id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": existingUser})
}
