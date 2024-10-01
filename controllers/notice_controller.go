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
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	noticeModel := models.Notice{
		Heading:    req.Heading,
		Content:    req.Content,
		Recipients: req.Recipients,
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
	if err := database.DB.Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}

func GetRecruiterNotice(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	var notices []models.Notice
	if err := database.DB.Where("? = ANY(recipients)", req.Email).Or("? = ANY(recipients)", "recruiter").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}

func GetStudentNotice(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	var notices []models.Notice
	if err := database.DB.Where("? = ANY(recipients)", req.Email).Or("? = ANY(recipients)", "student").Find(&notices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notices": notices})
}
