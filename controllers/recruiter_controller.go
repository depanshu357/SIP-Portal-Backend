package controllers

import (
	"fmt"
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

func UpdateRecruiterProfile(c *gin.Context) {
	var req struct {
		Company        string `json:"Company"`
		Email          string `json:"Email"`
		ContactNumber  string `json:"ContactNumber"`
		AdditionalInfo string `json:"AdditionalInfo"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	fmt.Println(req)
	if err := database.DB.Model(&existingUser).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}
