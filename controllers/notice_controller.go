package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"
	"sip/utils"

	"github.com/gin-gonic/gin"
)

func CreateNotice(c *gin.Context) {
	var req struct {
		Heading    string   `json:"heading" binding:"required"`
		Content    string   `json:"content" binding:"required"`
		Recipients []string `json:"recipients" binding:"required,min=1,dive,required"`
		Event      uint     `json:"events"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	noticeModel := models.Notice{
		Heading:    req.Heading,
		Content:    req.Content,
		Recipients: req.Recipients,
		Event:      req.Event,
	}
	if err := database.DB.Create(&noticeModel).Error; err != nil {
		utils.Logger.Sugar().Errorf("Failed to create Notice: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Notice"})
		return
	}
	// Further logic for handling the notice goes here
	c.JSON(http.StatusOK, gin.H{"message": "Notice created successfully"})
}

func GetAllNotice(c *gin.Context) {
	eventId := c.Query("event")
	if eventId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var notices []models.Notice
	if err := database.DB.Where("Event = ?", eventId).Order("created_at desc").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}

func GetRecruiterNotice(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var notices []models.Notice
	if err := database.DB.Where("? = ANY(recipients)", user_id).Or("? = ANY(recipients)", "recruiter").Order("created_at desc").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}

func GetStudentNotice(c *gin.Context) {
	eventId := c.Query("event")
	if eventId == "" {
		utils.Logger.Sugar().Error("Invalid Event ID")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var notices []models.Notice
	if err := database.DB.Where("? = ANY(recipients)", user_id).Or("? = ANY(recipients)", "student").Order("created_at desc").Find(&notices, "event = ?", eventId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}
