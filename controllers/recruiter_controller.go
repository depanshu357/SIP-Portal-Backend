package controllers

import (
	"fmt"
	"net/http"
	"sip/database"
	"sip/models"
	"time"

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

func GetJobDescriptions(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	eventID := c.Query("eventId")
	var existingUser models.Recruiter
	if err := database.DB.Where("user_id = ?", user_id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	var jobDescriptions []models.JobDescription
	if err := database.DB.Where("recruiter_id = ? AND event_id = ?", existingUser.ID, eventID).Find(&jobDescriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Job descriptions not found"})
		return
	}
	var jobDescriptionResponses []models.JobDescriptionResponse
	for _, jobDescription := range jobDescriptions {
		jobDescriptionResponses = append(jobDescriptionResponses, models.JobDescriptionResponse{
			ID:       jobDescription.ID,
			Title:    jobDescription.Title,
			Deadline: jobDescription.Deadline,
			Visible:  jobDescription.Visible,
		})
	}

	c.JSON(http.StatusOK, gin.H{"jobDescriptions": jobDescriptionResponses})
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
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		Stipend     string    `json:"stipend"`
		EventID     uint      `json:"eventId"`
		Eligibility []string  `json:"eligibility"`
		Deadline    time.Time `json:"deadline"`
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
		Deadline:    req.Deadline,
	}
	if err := database.DB.Create(&jobDescription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job description"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job Description created successfully"})
}

func EditJobDescription(c *gin.Context) {
	var req struct {
		ID          uint      `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Location    string    `json:"location"`
		Stipend     string    `json:"stipend"`
		EventID     uint      `json:"eventId"`
		Eligibility []string  `json:"eligibility"`
		Deadline    time.Time `json:"deadline"`
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
	var existingJobDescription models.JobDescription
	if err := database.DB.Where("id = ? AND recruiter_id = ?", req.ID, existingUser.ID).First(&existingJobDescription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Job description not found"})
		return
	}
	if err := database.DB.Model(&existingJobDescription).Updates(models.JobDescription{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Stipend:     req.Stipend,
		EventID:     req.EventID,
		Eligibility: req.Eligibility,
		Deadline:    req.Deadline,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job description"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job description updated successfully"})
}
