package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UserWithoutPassword struct {
	ID         uint `gorm:"primary_key"`
	Email      string
	CreatedAt  time.Time
	IsVerified bool
	Role       string
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

	// Map data from the original User struct to the new struct
	for _, user := range users {
		userWithoutPassword := UserWithoutPassword{
			ID:         user.ID,
			Email:      user.Email,
			CreatedAt:  user.CreatedAt,
			IsVerified: user.IsVerified,
			Role:       user.Role,
			// Map other fields as necessary
		}
		usersWithoutPassword = append(usersWithoutPassword, userWithoutPassword)
	}
	c.JSON(http.StatusOK, gin.H{"users": usersWithoutPassword})
}

func GetStudentList(c *gin.Context) {
	var students []models.Student
	if err := database.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": students})
}

func GetRecruiterList(c *gin.Context) {
	var reruiters []models.Recruiter
	if err := database.DB.Find(&reruiters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": reruiters})
}
