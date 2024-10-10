package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	academicYear := c.PostForm("academic_year")
	event := c.PostForm("event")
	if academicYear == "" || event == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required", "details": err.Error()})
		return
	}

	if err := c.SaveUploadedFile(file, "./uploads/"+academicYear+"/"+event+"/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": file.Filename})
}
