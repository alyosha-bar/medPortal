package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {

	roleSet := make(map[string]bool)
	for _, r := range allowedRoles {
		roleSet[r] = true
	}

	return func(c *gin.Context) {
		// get auth header
		authHeader := c.GetHeader("Authorization")

		// validate header
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// verify token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "could not extract claims"})
			return
		}

		role := claims["role"].(string)

		// check role
		if !roleSet[role] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized for this route"})
			return
		}

		// attach claims
		c.Set("user_id", claims["user_id"])
		c.Set("username", claims["username"])
		c.Set("role", role)

		// continue
		c.Next()
	}

}
