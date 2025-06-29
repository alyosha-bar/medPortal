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
	"golang.org/x/crypto/bcrypt"

	mock_handlers "github.com/alyosha-bar/medPortal/mocks"
)

func TestAuthHandler_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockAuthService(ctrl)
	handler := handlers.NewAuthHandler(mockService)

	// Prepare login input
	input := handlers.LoginInput{
		Username: "testuser",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(input)

	// Create a hashed password for comparison
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		ID:       1,
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     "user",
	}

	// Setup mock expectations
	mockService.EXPECT().GetUserByUsername(input.Username).Return(user, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Contains(t, resp, "token")
	assert.Contains(t, resp, "user")
	assert.Equal(t, float64(user.ID), resp["user"].(map[string]interface{})["id"])
	assert.Equal(t, user.Username, resp["user"].(map[string]interface{})["username"])
}

func TestAuthHandler_Login_BadRequest(t *testing.T) {
	handler := handlers.NewAuthHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// Send invalid JSON
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewReader([]byte(`invalid-json`)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.Login(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "username and password required")
}

func TestAuthHandler_Login_InvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockAuthService(ctrl)
	handler := handlers.NewAuthHandler(mockService)

	input := handlers.LoginInput{
		Username: "testuser",
		Password: "wrongpassword",
	}
	jsonBody, _ := json.Marshal(input)

	user := models.User{
		ID:       1,
		Username: input.Username,
		Password: "$2a$10$somethinghashed", // some hashed password not matching "wrongpassword"
		Role:     "user",
	}

	mockService.EXPECT().GetUserByUsername(input.Username).Return(user, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid credentials")
}

func TestAuthHandler_Login_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockAuthService(ctrl)
	handler := handlers.NewAuthHandler(mockService)

	input := handlers.LoginInput{
		Username: "unknownuser",
		Password: "password123",
	}
	jsonBody, _ := json.Marshal(input)

	mockService.EXPECT().GetUserByUsername(input.Username).Return(models.User{}, errors.New("not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "invalid credentials")
}

func TestAuthHandler_SignUp_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockAuthService(ctrl)
	handler := handlers.NewAuthHandler(mockService)

	newUser := models.User{
		Username: "newuser",
		Password: "password123",
		Role:     "user",
	}
	jsonBody, _ := json.Marshal(newUser)

	mockService.EXPECT().SignUp(gomock.Any()).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/signup", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.SignUp(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user created successfully")
}

func TestAuthHandler_SignUp_BadRequest(t *testing.T) {
	handler := handlers.NewAuthHandler(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/signup", bytes.NewReader([]byte(`invalid-json`)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

func TestAuthHandler_SignUp_MissingFields(t *testing.T) {
	handler := handlers.NewAuthHandler(nil)

	// Missing Role field
	newUser := models.User{
		Username: "newuser",
		Password: "password123",
		Role:     "",
	}
	jsonBody, _ := json.Marshal(newUser)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/signup", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "all fields required")
}

func TestAuthHandler_SignUp_HashPasswordError(t *testing.T) {
	// This is tricky to simulate since bcrypt.GenerateFromPassword rarely fails.
	// We can patch the bcrypt.GenerateFromPassword if using interfaces or wrappers.
	// For now, just skip this test or mark as TODO.
}

func TestAuthHandler_SignUp_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_handlers.NewMockAuthService(ctrl)
	handler := handlers.NewAuthHandler(mockService)

	newUser := models.User{
		Username: "newuser",
		Password: "password123",
		Role:     "user",
	}
	jsonBody, _ := json.Marshal(newUser)

	mockService.EXPECT().SignUp(gomock.Any()).Return(errors.New("service error"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/signup", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.SignUp(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to create user")
}
