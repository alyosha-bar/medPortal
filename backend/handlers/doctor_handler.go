package handlers

import (
	"net/http"

	"github.com/alyosha-bar/medPortal/services"
	"github.com/gin-gonic/gin"
)

func GetPatientsByDoctor(c *gin.Context) {

	// extract user id (doctor id)
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
	}

	// check type
	userID, ok := userIDVal.(uint)
	if !ok {
		if f, ok := userIDVal.(float64); ok {
			userID = uint(f)
		} else {
			c.JSON(500, gin.H{"error": "invalid user_id type"})
			return
		}
	}

	// call services to fetch patients
	patients, err := services.GetPatientsByDoctor(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
