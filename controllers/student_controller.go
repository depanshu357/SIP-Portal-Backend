package controllers

import (
	"net/http"
	"sip/database"
	"sip/models"

	"github.com/gin-gonic/gin"
)

func GetStudentProfile(c *gin.Context) {
	var id = c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID query parameter is required"})
		return
	}
	var existingUser models.Student
	if err := database.DB.Where("id = ?", id).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": existingUser})
}

func UpdateProfile(c *gin.Context) {
	var req struct {
		ID                     uint   `json:"id" binding:"required"`                  // id
		Name                   string `json:"name" binding:"required"`                // name
		RollNumber             string `json:"rollNumber"`                             // rollNumber
		Email                  string `json:"email" binding:"email"`                  // email
		Department             string `json:"department"`                             // department
		SecondaryDepartment    string `json:"secondaryDepartment"`                    // secondaryDepartment
		Specialisation         string `json:"specialisation"`                         // specialisation
		Gender                 string `json:"gender"`                                 // gender
		DOB                    string `json:"dob"`                                    // dob (you can use `time.Time` for date handling if needed)
		AlternateContactNumber string `json:"alternateContactNumber"`                 // alternateContactNumber
		CurrentCPI             string `json:"currentCPI"`                             // currentCPI
		TenthBoard             string `json:"tenthBoard"`                             // tenthBoard
		TenthMarks             string `json:"tenthMarks"`                             // tenthMarks
		TenthBoardYear         string `json:"tenthBoardYear"`                         // tenthBoardYear
		EntranceExam           string `json:"entranceExam"`                           // entranceExam
		Category               string `json:"category"`                               // category
		CurrentAddress         string `json:"currentAddress"`                         // currentAddress
		Disability             string `json:"disability"`                             // disability
		ExpectedGraduationYear string `json:"expectedGraduationYear"`                 // expectedGraduationYear
		Program                string `json:"program"`                                // program
		SecondaryProgram       string `json:"secondaryProgram"`                       // secondaryProgram
		Preference             string `json:"preference"`                             // preference
		PersonalEmail          string `json:"personalEmail"`                          // personalEmail
		ContactNumber          string `json:"contactNumber"`                          // contactNumber
		WhatsappNumber         string `json:"whatsappNumber"`                         // whatsappNumber
		TwelfthBoardYear       string `json:"twelfthBoardYear"`                       // twelfthBoardYear
		TwelfthBoard           string `json:"twelfthBoard"`                           // twelfthBoard
		TwelfthMarks           string `json:"twelfthMarks"`                           // twelfthMarks
		EntranceExamRank       string `json:"entranceExamRank"`                       // entranceExamRank
		CategoryRank           string `json:"categoryRank"`                           // categoryRank
		PermanentAddress       string `json:"permanentAddress"`                       // permanentAddress
		FriendsName            string `json:"friendsName"`                            // friendsName
		FriendsContactDetails  string `json:"friendsContactDetails"`                  // friendsContactDetails
		IsVerified             bool   `json:"isVerified" gorm:"default:false"`        // IsVerified
		IsProfileVerified      bool   `json:"isProfileVerified" gorm:"default:false"` // IsProfileVerified
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input"})
		return
	}
	var existingUser models.Student
	if err := database.DB.Where("id = ?", req.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	existingUser = models.Student{
		Name:                   req.Name,
		RollNumber:             req.RollNumber,
		Email:                  req.Email,
		Department:             req.Department,
		SecondaryDepartment:    req.SecondaryDepartment,
		Specialisation:         req.Specialisation,
		Gender:                 req.Gender,
		DOB:                    req.DOB,
		AlternateContactNumber: req.AlternateContactNumber,
		CurrentCPI:             req.CurrentCPI,
		TenthBoard:             req.TenthBoard,
		TenthMarks:             req.TenthMarks,
		TenthBoardYear:         req.TenthBoardYear,
		EntranceExam:           req.EntranceExam,
		Category:               req.Category,
		CurrentAddress:         req.CurrentAddress,
		Disability:             req.Disability,
		ExpectedGraduationYear: req.ExpectedGraduationYear,
		Program:                req.Program,
		SecondaryProgram:       req.SecondaryProgram,
		Preference:             req.Preference,
		PersonalEmail:          req.PersonalEmail,
		ContactNumber:          req.ContactNumber,
		WhatsappNumber:         req.WhatsappNumber,
		TwelfthBoardYear:       req.TwelfthBoardYear,
		TwelfthBoard:           req.TwelfthBoard,
		TwelfthMarks:           req.TwelfthMarks,
		EntranceExamRank:       req.EntranceExamRank,
		CategoryRank:           req.CategoryRank,
		PermanentAddress:       req.PermanentAddress,
		FriendsName:            req.FriendsName,
		FriendsContactDetails:  req.FriendsContactDetails,
		IsVerified:             req.IsVerified,
		IsProfileVerified:      req.IsProfileVerified,
	}
	if err := database.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
