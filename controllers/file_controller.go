package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	academicYear := c.PostForm("academic_year")
	event := c.PostForm("event")
	category := c.PostForm("category")
	if academicYear == "" || event == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required", "details": err.Error()})
		return
	}

	filePath := "./uploads/" + academicYear + "/" + event + "/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	user_id := uint(c.MustGet("user_id").(float64))
	fileModel := models.File{
		UserID:       user_id,
		Name:         file.Filename,
		Event:        event,
		Path:         filePath,
		IsVerified:   false,
		Category:     category,
		AcademicYear: academicYear,
	}
	if err := database.DB.Create(&fileModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": file.Filename})
}

func GetResumeList(c *gin.Context) {
	event := c.Query("event")
	user_id := uint(c.MustGet("user_id").(float64))
	var files []models.File
	if err := database.DB.Where("user_id = ? AND event = ?", user_id, event).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}
