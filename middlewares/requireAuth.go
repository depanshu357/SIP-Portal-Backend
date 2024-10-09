package middleware

import (
	"fmt"
	"net/http"
	"os"
	"sip/database"
	"sip/models"
	"sip/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	// ID         int    `json:"id"`
	// Email      string `json:"email"`
	// IsVerified bool   `json:"isVerified"`
	// Role       string `json:"role"`
	jwt.RegisteredClaims
	jwt.Token
}

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	// utils.Logger.Sugar().Info(c.Request.Header)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	utils.Logger.Sugar().Info(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println(err)
		utils.Logger.Sugar().Panic(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		subStr, ok := claims["sub"].(string)
		if !ok {
			utils.Logger.Sugar().Error("Invalid type for sub claim")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Parse the sub string into a UUID
		subUUID, err := uuid.Parse(subStr)
		if err != nil {
			utils.Logger.Sugar().Error("Invalid UUID format for sub claim: ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Query the database using the parsed UUID
		var user models.User
		if err := database.DB.First(&user, "id = ?", subUUID).Error; err != nil {
			utils.Logger.Sugar().Error("User not found: ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the user ID is a nil UUID, reject the request
		if user.ID == uuid.Nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
		c.Set("role", claims["role"])

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
