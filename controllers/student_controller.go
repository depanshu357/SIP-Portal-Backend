package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"

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
