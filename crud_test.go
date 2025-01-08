package main

import (
	"bytes"
	"crud_it_krasava/storage"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserHandlerWithMock(t *testing.T) {
	mockDAO := new(storage.MockUserStorage)
	handler = NewHandler(mockDAO)
	router := setupRouter()

	testUser := storage.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Age:       30,
	}

	createdUser := testUser
	createdUser.ID = uuid.New()
	createdUser.Created = testUser.Created

	mockDAO.On("CreateUser", mock.Anything).Return(createdUser, nil)

	req, w := makeRequest(http.MethodPost, "/users", testUser)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseUser storage.User
	_ = json.NewDecoder(w.Body).Decode(&responseUser)
	assert.Equal(t, createdUser.ID, responseUser.ID)
	assert.Equal(t, createdUser.Firstname, responseUser.Firstname)
	assert.Equal(t, createdUser.Lastname, responseUser.Lastname)
	assert.Equal(t, createdUser.Email, responseUser.Email)
	assert.Equal(t, createdUser.Age, responseUser.Age)

	mockDAO.AssertExpectations(t)
}

func TestGetUserHandlerWithMock(t *testing.T) {
	mockDAO := new(storage.MockUserStorage)
	handler = NewHandler(mockDAO)
	router := setupRouter()

	testID := uuid.New()
	testUser := storage.User{
		ID:        testID,
		Firstname: "Jane",
		Lastname:  "Doe",
		Email:     "jane.doe@example.com",
		Age:       25,
		Created:   time.Now(),
	}

	mockDAO.On("GetUserByID", testID).Return(testUser, nil)

	req, w := makeRequest(http.MethodGet, "/user/"+testID.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseUser storage.User
	_ = json.NewDecoder(w.Body).Decode(&responseUser)
	assert.Equal(t, testUser.ID, responseUser.ID)
	assert.Equal(t, testUser.Firstname, responseUser.Firstname)
	assert.Equal(t, testUser.Lastname, responseUser.Lastname)
	assert.Equal(t, testUser.Email, responseUser.Email)
	assert.Equal(t, testUser.Age, responseUser.Age)

	mockDAO.AssertExpectations(t)
}

func TestUpdateUserHandlerWithMock(t *testing.T) {
	mockDAO := new(storage.MockUserStorage)
	handler = NewHandler(mockDAO)
	router := setupRouter()

	testID := uuid.New()
	updatedUser := storage.User{
		Firstname: "Jacob",
		Lastname:  "Smith",
		Email:     "jacob.smith@example.com",
		Age:       45,
	}
	mockDAO.On("UpdateUser", testID, updatedUser).Return(updatedUser, nil)

	req, w := makeRequest(http.MethodPatch, "/user/"+testID.String(), updatedUser)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockDAO.AssertExpectations(t)
}

func makeRequest(method, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	var requestBody []byte
	if body != nil {
		requestBody, _ = json.Marshal(body)
	}
	request := httptest.NewRequest(method, url, bytes.NewReader(requestBody))
	response := httptest.NewRecorder()
	return request, response
}
