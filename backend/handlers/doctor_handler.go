package handlers

import (
	"net/http"
	"strconv"

	"github.com/alyosha-bar/medPortal/models"
	"github.com/gin-gonic/gin"
)

type DoctorService interface {
	GetPatientsByDoctor(doctorID uint) ([]models.Patient, error)
	UpdateMedicalNotes(doctorID uint, patientID uint, medicalNotes string) (models.Patient, error)
}

type DoctorHandler struct {
	Service DoctorService
}

func NewDoctorHandler(service DoctorService) *DoctorHandler {
	return &DoctorHandler{Service: service}
}

func (h *DoctorHandler) GetPatientsByDoctor(c *gin.Context) {

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
	patients, err := h.Service.GetPatientsByDoctor(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

type UpdateInput struct {
	MedicalNotes string `json:"medicalNotes"`
}

func (h *DoctorHandler) UpdateMedicalNotes(c *gin.Context) {
	// extract user id (doctor id)
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
		return
	}

	// check type
	var userID uint
	switch v := userIDVal.(type) {
	case uint:
		userID = v
	case float64:
		userID = uint(v)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id type"})
		return
	}

	var body UpdateInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "notes not provided."})
		return
	}

	patientIDStr := c.Param("patient_id")
	patientID64, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}
	patientID := uint(patientID64)

	// pass into services
	patient, err := h.Service.UpdateMedicalNotes(userID, patientID, body.MedicalNotes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}
