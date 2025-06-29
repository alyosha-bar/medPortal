package handlers

import (
	"net/http"

	"github.com/alyosha-bar/medPortal/helper"
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	// decode username & password
	var credentials LoginInput

	// bind JSON
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password required"})
		return
	}

	// validate credentials
	// fetch user
	var user models.User
	user, err := services.GetUserByUsername(credentials.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// issue JWT
	tokenString, err := helper.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate JWT"})
		return
	}

	// successful response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})

}

func SignUp(c *gin.Context) {
	// create a User object from request
	var newUser models.User

	// bind the JSON to user
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// basic validation
	if newUser.Username == "" || newUser.Password == "" || newUser.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all fields required"})
		return
	}

	// check for existing user
	// var existing models.User --> TODO

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// switch plaintext and hashed password
	newUser.Password = string(hashedPassword)

	// create new duser in DB (pass into services)
	err = services.SignUp(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, "user created successfully")
}
