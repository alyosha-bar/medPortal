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

func TestGetAllPatients_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	patients := []models.Patient{
		{ID: 1, Firstname: "Alice"},
		{ID: 2, Firstname: "Bob"},
	}

	mockService.EXPECT().GetAllPatients().Return(patients, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.GetAllPatients(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp handlers.ManyPatientResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.Data, 2)
	assert.Equal(t, "Alice", resp.Data[0].Firstname)
}

func TestGetAllPatients_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	mockService.EXPECT().GetAllPatients().Return(nil, errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.GetAllPatients(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to fetch patients")
}

func TestGetPatient_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	patient := models.Patient{ID: 1, Firstname: "Alice"}

	mockService.EXPECT().GetPatient(uint(1)).Return(patient, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}

	handler.GetPatient(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Patient
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", resp.Firstname)
}

func TestGetPatient_InvalidID(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "notanint"}}

	handler.GetPatient(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid patient id")
}

func TestGetPatient_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	mockService.EXPECT().GetPatient(uint(1)).Return(models.Patient{}, errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}

	handler.GetPatient(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to get patient details")
}

func TestRegisterPatient_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	newPatient := models.Patient{Firstname: "New", Lastname: "Patient"}
	jsonBody, _ := json.Marshal(newPatient)

	mockService.EXPECT().RegisterPatient(newPatient).Return(newPatient, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/patients", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.RegisterPatient(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Patient
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "New", resp.Firstname)
}

func TestRegisterPatient_BadRequest(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/patients", bytes.NewReader([]byte(`invalid-json`)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.RegisterPatient(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid patient data")
}

func TestRegisterPatient_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	newPatient := models.Patient{Firstname: "New", Lastname: "Patient"}
	jsonBody, _ := json.Marshal(newPatient)

	mockService.EXPECT().RegisterPatient(newPatient).Return(models.Patient{}, errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/patients", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.RegisterPatient(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to register patient")
}

func TestDeletePatientProfile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	mockService.EXPECT().DeletePatientProfile(uint(1)).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}

	handler.DeletePatientProfile(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Deleted Profile")
}

func TestDeletePatientProfile_InvalidID(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "notanint"}}

	handler.DeletePatientProfile(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid patient id")
}

func TestDeletePatientProfile_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	mockService.EXPECT().DeletePatientProfile(uint(1)).Return(errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}

	handler.DeletePatientProfile(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to delete patient profile")
}

func TestUpdateField_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	update := handlers.UpdateBody{
		Field: "firstname",
		Value: "UpdatedName",
	}
	jsonBody, _ := json.Marshal(update)

	updatedPatient := models.Patient{ID: 1, Firstname: "UpdatedName"}

	mockService.EXPECT().UpdateField(uint(1), "firstname", "UpdatedName").Return(updatedPatient, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("PATCH", "/patients/1", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateField(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Patient
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "UpdatedName", resp.Firstname)
}

func TestUpdateField_InvalidPatientID(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	jsonBody := []byte(`{"field":"firstname","value":"test"}`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "badID"}}
	c.Request = httptest.NewRequest("PATCH", "/patients/badID", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateField(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid patient ID")
}

func TestUpdateField_InvalidField(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	jsonBody := []byte(`{"field":"unknownfield","value":"test"}`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("PATCH", "/patients/1", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateField(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "field update not allowed")
}

func TestUpdateField_InvalidAgeValue(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	jsonBody := []byte(`{"field":"age","value":"notanint"}`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("PATCH", "/patients/1", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateField(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid age value")
}

func TestUpdateField_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	update := handlers.UpdateBody{
		Field: "firstname",
		Value: "UpdatedName",
	}
	jsonBody, _ := json.Marshal(update)

	mockService.EXPECT().UpdateField(uint(1), "firstname", "UpdatedName").Return(models.Patient{}, errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("PATCH", "/patients/1", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateField(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to update patient")
}

func TestGetAllDoctors_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	doctors := []models.User{
		{ID: 1, Username: "doc1"},
		{ID: 2, Username: "doc2"},
	}

	mockService.EXPECT().GetAllDoctors().Return(doctors, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.GetAllDoctors(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp []models.User
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "doc1", resp[0].Username)
}

func TestGetAllDoctors_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	mockService.EXPECT().GetAllDoctors().Return(nil, errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handler.GetAllDoctors(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to fetch doctors")
}

func TestAssignPatient_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	body := handlers.AssignBody{DoctorID: 10}
	jsonBody, _ := json.Marshal(body)

	patient := models.Patient{ID: 1, Firstname: "Alice"}

	mockService.EXPECT().AssignPatient(uint(1), uint(10)).Return(patient, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("POST", "/patients/1/assign", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.AssignPatient(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Patient
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", resp.Firstname)
}

func TestAssignPatient_InvalidPatientID(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	jsonBody := []byte(`{"doctorID": 10}`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "badID"}}
	c.Request = httptest.NewRequest("POST", "/patients/badID/assign", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.AssignPatient(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid patient ID")
}

func TestAssignPatient_BadRequest(t *testing.T) {
	handler := handlers.NewReceptionHandler(nil)

	jsonBody := []byte(`invalid-json`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("POST", "/patients/1/assign", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.AssignPatient(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid body")
}

func TestAssignPatient_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockReceptionService(ctrl)
	handler := handlers.NewReceptionHandler(mockService)

	body := handlers.AssignBody{DoctorID: 10}
	jsonBody, _ := json.Marshal(body)

	mockService.EXPECT().AssignPatient(uint(1), uint(10)).Return(models.Patient{}, errors.New("db error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "patient_id", Value: "1"}}
	c.Request = httptest.NewRequest("POST", "/patients/1/assign", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.AssignPatient(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to assign patient")
}
