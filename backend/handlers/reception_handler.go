package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/alyosha-bar/medPortal/models"
	"github.com/gin-gonic/gin"
)

type ManyPatientResponse struct {
	Data []models.Patient `json:"data"`
}

type OnePatientResponse struct {
	Data models.Patient `json:"data"`
}

type ReceptionService interface {
	GetAllPatients() ([]models.Patient, error)
	GetPatient(patientID uint) (models.Patient, error)
	RegisterPatient(patient models.Patient) (models.Patient, error)
	DeletePatientProfile(patientID uint) error
	UpdateField(patientID uint, field string, value interface{}) (models.Patient, error)
	GetAllDoctors() ([]models.User, error)
	AssignPatient(patientID uint, doctorID uint) (models.Patient, error)
}

type ReceptionHandler struct {
	Service ReceptionService
}

func NewReceptionHandler(service ReceptionService) *ReceptionHandler {
	return &ReceptionHandler{Service: service}
}

func (h *ReceptionHandler) GetAllPatients(c *gin.Context) {
	patients, err := h.Service.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch patients"})
		return
	}
	c.JSON(http.StatusOK, ManyPatientResponse{Data: patients})
}

func (h *ReceptionHandler) GetPatient(c *gin.Context) {
	patientIDStr := c.Param("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	patient, err := h.Service.GetPatient(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get patient details"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *ReceptionHandler) RegisterPatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient data"})
		return
	}

	patientReturn, err := h.Service.RegisterPatient(patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register patient"})
		return
	}

	c.JSON(http.StatusOK, patientReturn)
}

func (h *ReceptionHandler) DeletePatientProfile(c *gin.Context) {
	patientIDStr := c.Param("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	err = h.Service.DeletePatientProfile(uint(patientID))
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

func (h *ReceptionHandler) UpdateField(c *gin.Context) {
	patientIDStr := c.Param("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
		return
	}

	var body UpdateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input format"})
		return
	}

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

	var value interface{} = body.Value
	if field == "age" {
		ageInt, err := strconv.Atoi(body.Value)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age value"})
			return
		}
		value = ageInt
	}

	patient, err := h.Service.UpdateField(uint(patientID), field, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *ReceptionHandler) GetAllDoctors(c *gin.Context) {
	doctors, err := h.Service.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch doctors"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

type AssignBody struct {
	DoctorID uint `json:"doctorID"`
}

func (h *ReceptionHandler) AssignPatient(c *gin.Context) {
	patientIDStr := c.Param("patient_id")
	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
		return
	}

	var body AssignBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	patient, err := h.Service.AssignPatient(uint(patientID), body.DoctorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to assign patient"})
		return
	}

	c.JSON(http.StatusOK, patient)
}
