package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var req struct {
		Title        string    `json:"title" binding:"required"`
		StartDate    time.Time `json:"start_date" binding:"required"`
		AcademicYear string    `json:"academic_year" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	eventModel := models.Event{
		Title:        req.Title,
		StartDate:    req.StartDate,
		IsActive:     true,
		AcademicYear: req.AcademicYear,
	}
	if err := database.DB.Create(&eventModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully"})
}

func GetAllEvents(c *gin.Context) {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func GetPublicEvents(c *gin.Context) {
	var events []models.Event
	if err := database.DB.Where("is_active = ?", true).Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching events"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

func ToggleEventActivation(c *gin.Context) {
	var req struct {
		ID       uint `json:"id" binding:"required"`
		IsActive bool `json:"IsActive"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	var event models.Event
	if err := database.DB.Where("id = ?", req.ID).First(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	event.IsActive = req.IsActive
	if err := database.DB.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
