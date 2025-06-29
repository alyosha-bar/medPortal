package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alyosha-bar/medPortal/handlers"
	"github.com/alyosha-bar/medPortal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_handlers "github.com/alyosha-bar/medPortal/mocks"
)

func TestGetPatientsByDoctor_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)

	doctorID := uint(5)
	patients := []models.Patient{
		{ID: 1, Firstname: "John", Lastname: "Doe"},
		{ID: 2, Firstname: "Jane", Lastname: "Smith"},
	}

	mockService.
		EXPECT().
		GetPatientsByDoctor(doctorID).
		Return(patients, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// set user_id in context as uint
	c.Set("user_id", doctorID)

	handler.GetPatientsByDoctor(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []models.Patient
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, patients, resp)
}

func TestGetPatientsByDoctor_MissingUserID(t *testing.T) {
	handler := handlers.NewDoctorHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// not setting user_id to simulate missing user ID

	handler.GetPatientsByDoctor(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "missing user id")
}

func TestGetPatientsByDoctor_InvalidUserIDType(t *testing.T) {
	handler := handlers.NewDoctorHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set user_id with invalid type (string)
	c.Set("user_id", "not-a-number")

	handler.GetPatientsByDoctor(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "invalid user_id type")
}

func TestGetPatientsByDoctor_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)

	doctorID := uint(5)

	mockService.
		EXPECT().
		GetPatientsByDoctor(doctorID).
		Return(nil, errors.New("some service error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_id", doctorID)

	handler.GetPatientsByDoctor(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to fetch patients")
}

func TestUpdateMedicalNotes_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)

	doctorID := uint(5)
	patientID := uint(10)
	newNotes := "Updated notes"

	inputBody := handlers.UpdateInput{MedicalNotes: newNotes}
	jsonBody, _ := json.Marshal(inputBody)

	updatedPatient := models.Patient{
		ID:           patientID,
		MedicalNotes: newNotes,
	}

	mockService.
		EXPECT().
		UpdateMedicalNotes(doctorID, patientID, newNotes).
		Return(updatedPatient, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set user_id and patient_id param
	c.Set("user_id", doctorID)
	c.Params = gin.Params{{Key: "patient_id", Value: "10"}}
	c.Request = httptest.NewRequest("PUT", "/patients/10/notes", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateMedicalNotes(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Patient
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, updatedPatient, resp)
}

func TestUpdateMedicalNotes_MissingUserID(t *testing.T) {
	handler := handlers.NewDoctorHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// No user_id set
	handler.UpdateMedicalNotes(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "missing user id")
}

func TestUpdateMedicalNotes_InvalidUserIDType(t *testing.T) {
	handler := handlers.NewDoctorHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_id", "invalid")

	handler.UpdateMedicalNotes(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "invalid user_id type")
}

func TestUpdateMedicalNotes_InvalidPatientID(t *testing.T) {
	handler := handlers.NewDoctorHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_id", uint(1))
	c.Params = gin.Params{{Key: "patient_id", Value: "notanint"}}

	handler.UpdateMedicalNotes(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "notes not provided") // changed line
}

func TestUpdateMedicalNotes_BadRequestBody(t *testing.T) {
	handler := handlers.NewDoctorHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_id", uint(1))
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("PUT", "/patients/1/notes", bytes.NewReader([]byte(`invalid-json`)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateMedicalNotes(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "notes not provided")
}

func TestUpdateMedicalNotes_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockDoctorService(ctrl)
	handler := handlers.NewDoctorHandler(mockService)

	doctorID := uint(5)
	patientID := uint(10)
	newNotes := "Updated notes"

	mockService.
		EXPECT().
		UpdateMedicalNotes(doctorID, patientID, newNotes).
		Return(models.Patient{}, errors.New("service error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_id", doctorID)
	c.Params = gin.Params{{Key: "patient_id", Value: "10"}}
	inputBody := handlers.UpdateInput{MedicalNotes: newNotes}
	jsonBody, _ := json.Marshal(inputBody)
	c.Request = httptest.NewRequest("PUT", "/patients/10/notes", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateMedicalNotes(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to update patient")
}
