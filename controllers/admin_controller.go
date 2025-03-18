package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sip/database"
	"sip/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type UserWithoutPassword struct {
	ID             uint `gorm:"primary_key"`
	Email          string
	CreatedAt      time.Time
	HasAdminAccess bool
	Role           string
}

func GetAdminProfile(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var existingUser models.User
	if err := database.DB.Where("id = ?", user_id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": existingUser})

}

func GetAdminList(c *gin.Context) {
	var users []models.User
	if err := database.DB.Where("role = ?", "admin").Or("role = ?", "superadmin").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	var usersWithoutPassword []UserWithoutPassword

	for _, user := range users {
		userWithoutPassword := UserWithoutPassword{
			ID:             user.ID,
			Email:          user.Email,
			CreatedAt:      user.CreatedAt,
			HasAdminAccess: user.HasAdminAccess,
			Role:           user.Role,
		}
		usersWithoutPassword = append(usersWithoutPassword, userWithoutPassword)
	}
	c.JSON(http.StatusOK, gin.H{"users": usersWithoutPassword})
}

func GetStudentList(c *gin.Context) {
	type StudentInfo struct {
		ID                int    `gorm:"column:id"`
		RollNumber        string `gorm:"column:roll_number"`
		Email             string `gorm:"column:email"`
		IsProfileVerified bool   `gorm:"column:is_profile_verified"`
	}
	var students []StudentInfo
	if err := database.DB.Table("students").
		Joins("JOIN users ON users.id = students.user_id").Select("users.id as id", "roll_number", "students.email as email", "is_profile_verified").
		Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": students})
}

func GetStudentListForEvent(c *gin.Context) {
	type StudentInfo struct {
		ID                int             `gorm:"column:id"`
		RollNumber        string          `gorm:"column:roll_number"`
		Email             string          `gorm:"column:email"`
		VerifiedForEvents pq.Int64Array   `gorm:"type:integer[];default:'{}'"`
		FrozenForEvents   pq.Int64Array   `gorm:"type:integer[];default:'{}'"`
		ReasonForFreeze   json.RawMessage `gorm:"type:jsonb;default:'{}'::jsonb"`
	}
	var students []StudentInfo
	if err := database.DB.Table("students").
		Joins("JOIN users ON users.id = students.user_id").Select("users.id as id", "roll_number", "students.email as email", "verified_for_events", "frozen_for_events", "reason_for_freeze").
		Where("is_profile_verified = ?", true).
		Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": students})
}

func GetRecruiterList(c *gin.Context) {
	type RecruiterInfo struct {
		ID                int    `gorm:"column:id"`
		Company           string `gorm:"column:company"`
		Email             string `gorm:"column:email"`
		IsProfileVerified bool   `gorm:"column:is_profile_verified"`
	}
	var recruiters []RecruiterInfo
	if err := database.DB.Table("recruiters").
		Joins("JOIN users ON users.id = recruiters.user_id").Select("users.id as id", "company", "recruiters.email as email", "is_profile_verified").
		Find(&recruiters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": recruiters})
}

func GetRecruiterListForEvent(c *gin.Context) {
	type RecruiterInfo struct {
		ID                int             `gorm:"column:id"`
		Company           string          `gorm:"column:company"`
		Email             string          `gorm:"column:email"`
		ContactNumber     string          `gorm:"column:contact_number"`
		VerifiedForEvents pq.Int64Array   `gorm:"type:integer[];default:'{}'"`
		FrozenForEvents   pq.Int64Array   `gorm:"type:integer[];default:'{}'"`
		ReasonForFreeze   json.RawMessage `gorm:"type:jsonb;default:'{}'::jsonb"`
	}
	var recruiters []RecruiterInfo
	if err := database.DB.Table("recruiters").
		Joins("JOIN users ON users.id = recruiters.user_id").Select("users.id as id", "company", "recruiters.email as email", "verified_for_events", "frozen_for_events", "reason_for_freeze").
		Where("is_profile_verified = ?", true).
		Find(&recruiters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": recruiters})
}

type JobDescriptionList struct {
	ID       int       `gorm:"column:id"`
	Title    string    `gorm:"column:title"`
	Deadline time.Time `gorm:"column:deadline"`
	Visible  bool      `gorm:"column:visible"`
	Company  string    `gorm:"column:company"`
}

func GetAllJobDescriptions(c *gin.Context) {
	var jobDescriptionList []JobDescriptionList
	eventId := c.Query("event")
	if err := database.DB.Table("job_descriptions").
		Joins("JOIN recruiters ON recruiters.id = job_descriptions.recruiter_id").
		Select("job_descriptions.id, job_descriptions.title, job_descriptions.deadline, job_descriptions.visible, recruiters.company as company").
		Where("job_descriptions.event_id = ?", eventId).
		Find(&jobDescriptionList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobDescriptionList": jobDescriptionList})
}

func ChangeAdminAccess(c *gin.Context) {
	// user_id, exists := c.Get("user_id")
	var req struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	role, role_exists := c.Get("role")
	if !role_exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	if role != "superadmin" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not a Super Admin"})
		return
	}
	var existingUser models.User
	if err := database.DB.Where("id = ?", req.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	existingUser.HasAdminAccess = !existingUser.HasAdminAccess
	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Admin Access"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin Access updated successfully"})
}

func ChangeProfileVerification(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var existingUser models.User
	if err := database.DB.Where("id = ?", req.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	existingUser.IsProfileVerified = !existingUser.IsProfileVerified
	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Profile Verification"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile Verification updated successfully"})
}

func ToggleVerificationForEvent(c *gin.Context) {
	var req struct {
		ID    int `json:"id"`
		Event int `json:"event"`
	}
	fmt.Println(req)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var existingUser models.User
	if err := database.DB.Where("id = ?", req.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	eventExists := false
	for i, event := range existingUser.VerifiedForEvents {
		if event == int64(req.Event) {
			existingUser.VerifiedForEvents = append(existingUser.VerifiedForEvents[:i], existingUser.VerifiedForEvents[i+1:]...)
			eventExists = true
			break
		}
	}
	if !eventExists {
		existingUser.VerifiedForEvents = append(existingUser.VerifiedForEvents, int64(req.Event))
	}
	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify for event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated event verification"})
}

func ToggleFreezingForEvent(c *gin.Context) {
	var req struct {
		ID              int             `json:"id"`
		Event           int             `json:"event"`
		ReasonForFreeze json.RawMessage `json:"reasonForFreeze"`
	}
	fmt.Println(req)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var existingUser models.User
	if err := database.DB.Where("id = ?", req.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	eventExists := false
	for i, event := range existingUser.FrozenForEvents {
		if event == int64(req.Event) {
			existingUser.FrozenForEvents = append(existingUser.FrozenForEvents[:i], existingUser.FrozenForEvents[i+1:]...)
			eventExists = true
			break
		}
	}
	if !eventExists {
		existingUser.FrozenForEvents = append(existingUser.FrozenForEvents, int64(req.Event))
	}
	existingUser.ReasonForFreeze = req.ReasonForFreeze
	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify for event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated user event frozen status"})
}
