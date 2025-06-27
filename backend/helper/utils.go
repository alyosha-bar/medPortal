package helper

import (
	"os"
	"time"

	"github.com/alyosha-bar/medPortal/models"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// helper function
func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
