package controllers

import (
	"errors"
	"net/http"
	"os"
	"sip/database"
	"sip/models"
	"sip/services"
	"sip/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=6"`
		Role        string `json:"role" binding:"required"`
		RollNo      string `json:"rollNo"`
		CompanyName string `json:"companyName"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Sugar().Errorf("Signup validation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var existingUser models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.Logger.Sugar().Warnf("Signup failed: user %s already exists", req.Email)
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Sugar().Errorf("Password hashing failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.Logger.Sugar().Errorf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	if req.Role == "student" {
		student := models.Student{
			User:       user,
			Email:      req.Email,
			RollNumber: req.RollNo,
		}
		if err := database.DB.Create(&student).Error; err != nil {
			utils.Logger.Sugar().Errorf("Failed to create student: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
			return
		}
	} else if req.Role == "recruiter" {
		company := models.Recruiter{
			User:    user,
			Email:   req.Email,
			Company: req.CompanyName,
		}
		if err := database.DB.Create(&company).Error; err != nil {
			utils.Logger.Sugar().Errorf("Failed to create company: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
			return
		}
	} else if req.Role == "admin" {
		admin := models.Admin{
			User:  user,
			Email: req.Email,
		}
		if err := database.DB.Create(&admin).Error; err != nil {
			utils.Logger.Sugar().Errorf("Failed to create admin: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
			return
		}
	}

	var createdUser models.User
	database.DB.Where("email = ?", req.Email).First(&createdUser)
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.Logger.Sugar().Warnf("User not found: %s", req.Email)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	token, err := generateJWT(user.ID, user.Role)
	if err != nil {
		utils.Logger.Sugar().Errorf("Failed to generate JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)

	utils.Logger.Sugar().Infof("User created and logged in: %s", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "userInfo": createdUser})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Sugar().Errorf("Login validation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.Logger.Sugar().Warnf("User not found: %s", req.Email)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Logger.Sugar().Warnf("Invalid login attempt for user: %s", req.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if (user.Role == "admin" || user.Role == "superadmin") && !user.HasAdminAccess {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Don't have admin access"})
	}
	token, err := generateJWT(user.ID, user.Role)
	if err != nil {
		utils.Logger.Sugar().Errorf("Failed to generate JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", token, 3600*24, "/", "sipp.iitk.ac.in", true, true)

	utils.Logger.Sugar().Infof("User logged in: %s", user.Email)
	// println(token)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user, "token": token})
}

func generateJWT(userID uint, userRole string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"role": userRole,
		"exp":  time.Now().Add(30 * 24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func Validate(c *gin.Context) {
	user_id, _ := c.Get("user_id")

	// has a middleware that sets the user in the context

	c.JSON(http.StatusOK, gin.H{
		"message": user_id,
	})
}

func GenerateAndSendOTP(c *gin.Context) {
	var body struct {
		Email string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Delete expired OTPs
	database.DB.Where("deletion_time < ?", time.Now()).Delete(&models.Otp{})

	// Check if the email already exists in the database
	var existingUser models.User
	if err := database.DB.Where("email = ?", body.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email ID already exists"})
		return
	}

	// Check if an OTP has already been sent for this email and is still valid
	var existingOtp models.Otp
	result := database.DB.Where("email = ? AND deletion_time > ?", body.Email, time.Now()).First(&existingOtp)

	// If an OTP exists and is not expired, return an error indicating OTP was already sent
	if result.RowsAffected != 0 {
		c.JSON(http.StatusAlreadyReported, gin.H{"message": "OTP already sent"})
		return
	}

	// Generate a new OTP
	otp_size := 6
	otp, err := services.GenerateOTP(otp_size)
	if err != nil {
		utils.Logger.Sugar().Errorf("Failed to generate OTP: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	// Send OTP via email
	err = services.SendMail(body.Email, otp)
	if err != nil {
		utils.Logger.Sugar().Errorf("Failed to send OTP: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	// Create a new OTP record in the database
	otpModel := models.Otp{
		Email:        body.Email,
		Otp:          otp,
		DeletionTime: time.Now().Add(time.Minute * 10), // OTP expires after 10 minutes
	}

	// Insert the OTP record into the database
	if err := database.DB.Create(&otpModel).Error; err != nil {
		utils.Logger.Sugar().Errorf("Failed to create OTP record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store OTP"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func ChangePassword(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Sugar().Errorf("Change password validation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	//update password
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.Logger.Sugar().Warnf("User not found: %s", req.Email)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Sugar().Errorf("Password hashing failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	user.Password = string(hashedPassword)
	if err := database.DB.Save(&user).Error; err != nil {
		utils.Logger.Sugar().Errorf("Failed to update password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})

}

func VerifyOTP(c *gin.Context) {
	var body struct {
		Email string
		Otp   string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// check if the otp matches
	var user_otp models.Otp
	record := database.DB.First(&user_otp, "email = ?", body.Email)

	// Handle the case when no record is found
	if record.Error != nil {
		if errors.Is(record.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		// Handle any other database error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	// If OTP does not match, return an error
	if user_otp.Otp != body.Otp {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP does not match"})
		return
	}

	// If OTP matches, delete the OTP record from the database
	database.DB.Delete(&models.Otp{}, "email = ?", body.Email)

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
