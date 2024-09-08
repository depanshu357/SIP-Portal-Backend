package controllers

import (
	"log"
	"net/http"
	"os"
	"sip/database"
	"sip/models"
	"sip/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email      string
		Password   string
		IsVerified bool
		Role       string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	var checkUser models.User
	database.DB.Find(&checkUser, "email = ?", body.Email)
	if checkUser.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash), IsVerified: body.IsVerified, Role: body.Role}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// otp := "324353"
	// sendMail(user.Email, otp)

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
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
		log.Fatalf("Failed to generate OTP: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	// Send OTP via email
	err = services.SendMail(body.Email, otp)
	if err != nil {
		log.Fatalf("Failed to send OTP: %v", err)
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
		log.Fatalf("Failed to create OTP record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store OTP"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
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
	// check if the otp mathces
	var user_otp models.Otp
	record := database.DB.First(&user_otp, "email = ?", body.Email)
	// if otp does not match, return an error
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OTP does not match"})
		return
	}
	if user_otp.Otp != body.Otp {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OTP does not match"})
		return
	}
	// if otp matches, delete the otp record from the database
	database.DB.Delete(&models.Otp{}, "email = ?", body.Email)

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
