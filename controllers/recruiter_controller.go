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
		Company          string `json:"Company"`
		Email            string `json:"Email"`
		ContactNumber    string `json:"ContactNumber"`
		AdditionalInfo   string `json:"AdditionalInfo"`
		NatureOfBusiness string `json:"NatureOfBusiness"`
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

func CreateJobDescription(c *gin.Context) {
	var req struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Location    string   `json:"location"`
		Stipend     string   `json:"stipend"`
		EventID     uint     `json:"eventId"`
		Eligibility []string `json:"eligibility"`
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
	var existingEvent models.Event
	if err := database.DB.Where("id = ?", req.EventID).First(&existingEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	jobDescription := models.JobDescription{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Stipend:     req.Stipend,
		RecruiterID: existingUser.ID,
		EventID:     existingEvent.ID,
		Eligibility: req.Eligibility,
	}
	if err := database.DB.Create(&jobDescription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job description"})
		return
	}

	fmt.Println(req)
}
