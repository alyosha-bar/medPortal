package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/services"
	"github.com/gin-gonic/gin"
)

type ManyPatientResponse struct {
	Data []models.Patient `json:"data"`
}

type OnePatientResponse struct {
	Data models.Patient `json:"data"`
}

// GetAllPatients
// @Summary Gets All Patients
// @Description Lists all patients
// @Tags Receptionist
// @Success 200 {object} ManyPatientResponse
// @Router /api/v1/receptionist/patients [get]
func GetAllPatients(c *gin.Context) {

	patients, err := services.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, ManyPatientResponse{Data: patients})
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
	Field string `json:"field" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func UpdateField(c *gin.Context) {
	// Parse patient ID from URL parameters
	patientIDStr := c.Param("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
		return
	}

	// bind body
	var body UpdateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input format"})
		return
	}

	// validate allowed fields
	allowedFields := map[string]bool{
		"firstname":     true,
		"lastname":      true,
		"age":           true,
		"gender":        true,
		"medical_notes": true,
	}
	field := strings.ToLower(body.Field)
	if !allowedFields[field] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field update not allowed"})
		return
	}

	// handle type conversion
	var value interface{} = body.Value
	if field == "age" {
		ageInt, err := strconv.Atoi(body.Value)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age value"})
			return
		}
		value = ageInt
	}

	patient, err := services.UpdateField(uint(patientID), field, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func GetAllDoctors(c *gin.Context) {
	doctors, err := services.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch doctors"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

type AssignBody struct {
	DoctorID uint `json:"doctorID"`
}

func AssignPatient(c *gin.Context) {
	// Extract patientID from URL parameters
	patientIDStr := c.Param("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
		return
	}

	// extract body
	var body AssignBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	patient, err := services.AssignPatient(uint(patientID), uint(body.DoctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get patient details"})
		return
	}

	c.JSON(http.StatusOK, patient)
}
