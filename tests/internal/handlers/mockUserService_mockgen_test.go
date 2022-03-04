package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RomaBilka/BloGo/tests/internal/models"
	"github.com/RomaBilka/BloGo/tests/mockgen-mocks/mock_handlers"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAPI(t *testing.T) {
	testCases := []struct {
		name          string
		body          createUserRequest
		method        string
		create        func(mockUserService *mock_handlers.MockuserService)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "Ok",
			body:   userTest,
			method: http.MethodPost,
			create: func(mockUserService *mock_handlers.MockuserService) {
				mockUserService.EXPECT().CreateUser(models.User{Name: userTest.Name, Email: userTest.Email, Phone: userTest.Phone}).Return(models.User{Id: 1, Name: userTest.Name, Email: userTest.Email, Phone: userTest.Phone}, nil)
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
			create: func(mockUserService *mock_handlers.MockuserService) {
				mockUserService.EXPECT().CreateUser(models.User{}).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)

				type errorResponse struct {
					Error string `json:"error"`
				}
				response := &errorResponse{}

				err := json.NewDecoder(recorder.Body).Decode(response)

				assert.NoError(t, err)
				assert.Equal(t, response.Error, "Bad request, short user name")
			},
		},
		{
			name:   "StatusMethodNotAllowed",
			body:   userTest,
			method: http.MethodGet,
			create: func(mockUserService *mock_handlers.MockuserService) {
				mockUserService.EXPECT().CreateUser(models.User{}).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserService := mock_handlers.NewMockuserService(ctrl)
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