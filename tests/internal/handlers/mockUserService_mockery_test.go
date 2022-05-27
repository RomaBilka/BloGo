package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RomaBilka/BloGo/tests/internal/models"
	"github.com/RomaBilka/BloGo/tests/mockery-mocks/mock_handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserAPIV2(t *testing.T) {
	testCases := []struct {
		name          string
		body          createUserRequest
		method        string
		create        func(mockUserService *mock_handlers.UserService)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "Ok",
			body:   userTest,
			method: http.MethodPost,
			create: func(mockUserService *mock_handlers.UserService) {
				mockUserService.On("CreateUser", mock.Anything).Return(models.User{Id: 1, Name: userTest.Name, Email: userTest.Email, Phone: userTest.Phone}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)

				userResponse := &createUserResponse{}
				err := json.NewDecoder(recorder.Body).Decode(userResponse)
				assert.NoError(t, err)
				assert.Equal(t, userTest.Name, userResponse.Name)
				assert.Equal(t, uint64(1), userResponse.Id)
			},
		},
		{
			name:   "ShortName",
			body:   createUserRequest{"", test.Email, test.Phone},
			method: http.MethodPost,
			create: func(mockUserService *mock_handlers.UserService) {
				mockUserService.On("CreateUser", mock.Anything).Return(mock.Anything, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)

				type errorResponse struct {
					Error string `json:"error"`
				}
				response := &errorResponse{}

				err := json.NewDecoder(recorder.Body).Decode(response)
				assert.NoError(t, err)
				assert.Equal(t, "Bad request, short user name", response.Error)
			},
		},
		{
			name:   "StatusMethodNotAllowed",
			body:   userTest,
			method: http.MethodGet,
			create: func(mockUserService *mock_handlers.UserService) {
				mockUserService.On("CreateUser", mock.Anything).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {

			mockUserService := new(mock_handlers.UserService)
			testCase.create(mockUserService)
			uHttp := NewUserHttp(mockUserService)

			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(testCase.body)
			assert.NoError(t, err)

			req, err := http.NewRequest(testCase.method, "/create-user", b)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(uHttp.CreateUser)
			handler.ServeHTTP(recorder, req)

			testCase.checkResponse(recorder)
		})
	}
}
