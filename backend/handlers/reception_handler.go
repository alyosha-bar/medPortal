package handlers

import (
	"net/http"
	"strconv"

	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/services"
	"github.com/gin-gonic/gin"
)

func GetAllPatients(c *gin.Context) {

	patients, err := services.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func GetPatient(c *gin.Context) {
	patientIDStr := c.Param("patient_id")

	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	patient, err := services.GetPatient(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get patient details"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func RegisterPatient(c *gin.Context) {
	// bind JSON to patient object
	var patient models.Patient

	err := c.ShouldBindJSON(&patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient data"})
		return
	}

	patient_return, err := services.RegisterPatient(patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register patient"})
		return
	}

	c.JSON(http.StatusOK, patient_return)
}

func DeletePatientProfile(c *gin.Context) {
	patientIDStr := c.Param("patient_id")

	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	err = services.DeletePatientProfile(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete patient profile"})
		return
	}

	c.JSON(http.StatusOK, "Deleted Profile")
}

type UpdateBody struct {
	field string `json:"field" binding:"required"`
	value string `json:"value" binding:"required"`
}

func UpdateField(c *gin.Context) {
	// pull out body

	// pass in field and new value

	// return updated patient entity
}
