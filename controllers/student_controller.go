package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStudentProfile(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var existingUser models.Student
	if err := database.DB.Where("user_id = ?", user_id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": existingUser})
}

func GetStudentInfoForResumeName(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var existingUser models.Student
	if err := database.DB.Where("user_id = ?", user_id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": existingUser.Name, "rollNumber": existingUser.RollNumber, "program": existingUser.Program, "department": existingUser.Department})
}

func UpdateProfile(c *gin.Context) {
	var req struct {
		Name                   string `json:"name"`                   // name
		RollNumber             string `json:"rollNumber"`             // rollNumber
		Email                  string `json:"email"`                  // email
		Department             string `json:"department"`             // department
		SecondaryDepartment    string `json:"secondaryDepartment"`    // secondaryDepartment
		Specialisation         string `json:"specialisation"`         // specialisation
		Gender                 string `json:"gender"`                 // gender
		DOB                    string `json:"dob"`                    // dob (you can use `time.Time` for date handling if needed)
		AlternateContactNumber string `json:"alternateContactNumber"` // alternateContactNumber
		CurrentCPI             string `json:"currentCPI"`             // currentCPI
		TenthBoard             string `json:"tenthBoard"`             // tenthBoard
		TenthMarks             string `json:"tenthMarks"`             // tenthMarks
		TenthBoardYear         string `json:"tenthBoardYear"`         // tenthBoardYear
		EntranceExam           string `json:"entranceExam"`           // entranceExam
		Category               string `json:"category"`               // category
		CurrentAddress         string `json:"currentAddress"`         // currentAddress
		Disability             string `json:"disability"`             // disability
		ExpectedGraduationYear string `json:"expectedGraduationYear"` // expectedGraduationYear
		Program                string `json:"program"`                // program
		SecondaryProgram       string `json:"secondaryProgram"`       // secondaryProgram
		Preference             string `json:"preference"`             // preference
		PersonalEmail          string `json:"personalEmail"`          // personalEmail
		ContactNumber          string `json:"contactNumber"`          // contactNumber
		WhatsappNumber         string `json:"whatsappNumber"`         // whatsappNumber
		TwelfthBoardYear       string `json:"twelfthBoardYear"`       // twelfthBoardYear
		TwelfthBoard           string `json:"twelfthBoard"`           // twelfthBoard
		TwelfthMarks           string `json:"twelfthMarks"`           // twelfthMarks
		EntranceExamRank       string `json:"entranceExamRank"`       // entranceExamRank
		CategoryRank           string `json:"categoryRank"`           // categoryRank
		PermanentAddress       string `json:"permanentAddress"`       // permanentAddress
		FriendsName            string `json:"friendsName"`            // friendsName
		FriendsContactDetails  string `json:"friendsContactDetails"`  // friendsContactDetails
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input data"})
		return
	}
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var existingUser models.Student
	if err := database.DB.Where("user_id = ?", user_id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	if err := database.DB.Model(&existingUser).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func GetJobDescriptionListForStudent(c *gin.Context) {
	eventID := c.Query("eventID")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eventID is required"})
		return
	}
	type JobDescriptionList struct {
		ID       int       `gorm:"column:id"`
		Title    string    `gorm:"column:title"`
		Deadline time.Time `gorm:"column:deadline"`
		Company  string    `gorm:"column:company"`
	}
	var jobDescriptionList []JobDescriptionList
	if err := database.DB.Table("job_descriptions").
		Joins("JOIN recruiters ON recruiters.id = job_descriptions.recruiter_id").
		Select("job_descriptions.id, job_descriptions.title, job_descriptions.deadline, job_descriptions.visible, recruiters.company as company").
		Where("job_descriptions.event_id = ?", eventID).
		Where("job_descriptions.visible = ?", true).
		Find(&jobDescriptionList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jobDescriptionList": jobDescriptionList})
}
