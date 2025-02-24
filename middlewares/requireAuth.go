package middleware

import (
	"fmt"
	"net/http"
	"os"
	"sip/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	utils.Logger.Sugar().Info(c.Request.Header)
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

		c.Set("user_id", claims["sub"])
		c.Set("role", claims["role"])
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
