package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"
	"sip/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

func CreateNotice(c *gin.Context) {
	var req struct {
		Heading    string    `json:"heading" binding:"required"`
		Content    string    `json:"content" binding:"required"`
		Recipients []string  `json:"recipients" binding:"required,min=1,dive,required"`
		Event      uuid.UUID `json:"events"`
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
	var notices []models.Notice
	if err := database.DB.Order("created_at desc").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}

func GetRecruiterNotice(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email query parameter is required"})
		return
	}
	var notices []models.Notice
	if err := database.DB.Where("? = ANY(recipients)", email).Or("? = ANY(recipients)", "recruiter").Order("created_at desc").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}

func GetStudentNotice(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email query parameter is required"})
		return
	}
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
	// 	return
	// }
	var notices []models.Notice
	if err := database.DB.Where("? = ANY(recipients)", email).Or("? = ANY(recipients)", "student").Order("created_at desc").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}
