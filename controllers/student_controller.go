package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"

	"github.com/gin-gonic/gin"
)

func GetStudentProfile(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input"})
		return
	}
	var existingUser models.Student
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": existingUser})
}
