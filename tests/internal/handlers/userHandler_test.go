package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RomaBilka/BloGo/tests/internal/models"
	"github.com/stretchr/testify/assert"
)

func createUser(t *testing.T, user createUserRequest) *httptest.ResponseRecorder {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(user)

	req, err := http.NewRequest(http.MethodPost, "/create-user", b)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(uHttp.CreateUser)
	handler.ServeHTTP(recorder, req)

	return recorder
}

func TestCreateUser(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	recorder := createUser(t, userTest)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	userResponse := &createUserResponse{}
	err := json.NewDecoder(recorder.Body).Decode(userResponse)

	assert.NoError(t, err)

	userInDb, err := testUserRepository.GetUserById(models.UserId(userResponse.Id))

	assert.NoError(t, err)
	assert.Equal(t, userTest.Name, userInDb.Name)
	assert.Equal(t, userTest.Email, userInDb.Email)
	assert.Equal(t, userTest.Phone, userInDb.Phone)
}

func TestCreateSecondUser(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	createTestUser(t)
	recorder := createUser(t, userTest)

	errorResponse := &errorResponse{}
	err := json.NewDecoder(recorder.Body).Decode(errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, errorResponse)
	assert.Equal(t, "User with that phone or email already exists", errorResponse.Error)
}
