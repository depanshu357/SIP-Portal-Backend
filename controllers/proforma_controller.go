package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Proforma struct {
	ID          int            `gorm:"column:id"`
	Title       string         `gorm:"column:title"`
	Deadline    time.Time      `gorm:"column:deadline"`
	Visible     bool           `gorm:"column:visible"`
	Company     string         `gorm:"column:company"`
	Description string         `gorm:"column:description"`
	Stipend     string         `gorm:"column:stipend"`
	Location    string         `gorm:"column:location"`
	Eligibility pq.StringArray `gorm:"column:eligibility;type:text[]"`
}

func GetProforma(c *gin.Context) {
	proformaId := c.Query("proformaId")

	var proforma Proforma
	if err := database.DB.Table("job_descriptions").
		Joins("JOIN recruiters ON recruiters.id = job_descriptions.recruiter_id").
		Select("job_descriptions.id, job_descriptions.title, job_descriptions.deadline, job_descriptions.visible, job_descriptions.stipend, job_descriptions.location, job_descriptions.description, job_descriptions.eligibility, recruiters.company as company").
		Where("job_descriptions.id = ?", proformaId).
		First(&proforma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching proforma"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"proforma": proforma})
}

func ToggleProformaVisibility(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	var proforma models.JobDescription
	if err := database.DB.Where("id = ?", req.ID).First(&proforma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Proforma not found"})
		return
	}
	proforma.Visible = !proforma.Visible
	if err := database.DB.Save(&proforma).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update proforma"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Visibility status updated successfully"})
}

func GetListOfAppliedCandidates(c *gin.Context) {
	proformaId := c.Query("proformaId")
	type ApplicantInfo struct {
		ID            int    `gorm:"column:id"`
		FileID        uint   `gorm:"column:file_id"`
		Name          string `gorm:"column:name"`
		Email         string `gorm:"column:email"`
		ContactNumber string `gorm:"column:contact_number"`
		RollNumber    int    `gorm:"column:roll_number"`
	}
	var applicants []ApplicantInfo
	if err := database.DB.Table("applicants").
		Joins("JOIN students ON students.id = applicants.student_id").
		Select("applicants.id, applicants.file_id, students.name as name, students.email as email, students.contact_number as contact_number, students.roll_number as roll_number").
		Where("applicants.job_description_id = ?", proformaId).
		Find(&applicants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching proforma"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"applicants": applicants})
}
