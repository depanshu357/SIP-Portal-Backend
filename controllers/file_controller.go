package controllers

import (
	"net/http"
	"path"
	"sip/database"
	"sip/models"
	"sip/utils"
	"strings"

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

	fileName := strings.TrimSuffix(file.Filename, path.Ext(file.Filename))
	filePath := "./uploads/" + academicYear + "/" + event + "/" + fileName + path.Ext(file.Filename)
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
	academic_year := c.Query("academic_year")
	utils.Logger.Sugar().Info(event, academic_year)
	user_id := uint(c.MustGet("user_id").(float64))
	var files []models.File
	if err := database.DB.Where("user_id = ? AND event = ? AND academic_year = ?", user_id, event, academic_year).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}
