package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sip/database"
	"sip/models"
	"sip/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	eventID := c.PostForm("eventId")
	category := c.PostForm("category")
	fmt.Println("is it working?")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required", "details": err.Error()})
		return
	}
	var existingEvent models.Event
	if err := database.DB.Where("id = ?", eventID).First(&existingEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Proforma not found"})
		return
	}

	fileName := strings.TrimSuffix(file.Filename, path.Ext(file.Filename))
	filePath := "./uploads/" + existingEvent.AcademicYear + "/" + existingEvent.Title + "/" + fileName + path.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	user_id := uint(c.MustGet("user_id").(float64))
	fileModel := models.File{
		UserID:     user_id,
		Name:       file.Filename,
		EventID:    existingEvent.ID,
		Path:       filePath,
		IsVerified: false,
		Category:   category,
	}
	if err := database.DB.Create(&fileModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": file.Filename})
}

func GetResumeList(c *gin.Context) {
	eventID := c.Query("eventId")
	utils.Logger.Sugar().Info(eventID)
	user_id := uint(c.MustGet("user_id").(float64))
	var files []models.File
	if err := database.DB.Where("user_id = ? AND event_id = ?", user_id, eventID).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func GetResumeListForAdmin(c *gin.Context) {
	eventID := c.Query("eventId")
	utils.Logger.Sugar().Info(eventID)
	var files []models.File
	if err := database.DB.Where("event_id = ?", eventID).Find(&files).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func VerifyResume(c *gin.Context) {
	var req struct {
		ID    uint `json:"id"`
		Value bool `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := database.DB.Model(&models.File{}).Where("id = ?", req.ID).Update("is_verified", req.Value).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change verification status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File verification status changed"})
}

func DownloadFile(c *gin.Context) {
	id := c.Query("id")
	var file models.File
	if err := database.DB.Where("id = ?", id).First(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
		return
	}
	filePath := file.Path // File to download

	// Set the headers for file download
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", getFileSize(filePath)))

	// Send the file to the client
	c.File(filePath)
}

func DeleteFile(c *gin.Context) {
	id := c.Query("id")
	var file models.File
	if err := database.DB.Where("id = ?", id).First(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File not found"})
		return
	}
	filePath := file.Path
	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}
	if err := database.DB.Delete(&file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
func getFileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
