package handlers

import (
	"authentication/internal/models"
	"authentication/internal/models/mocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFindOne(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		email := "test@test.com"

		var mockedUser *models.User
		mockedUser = &models.User{
			Email:     email,
			Username:  "testUser",
			FirstName: "Test",
			LastName:  "User",
		}

		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("FindOne", mock.Anything).Return(mockedUser, nil)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"user": mockedUser,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "http://localhost:"+os.Getenv("PORT")+"/api/?email="+email, nil)
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())

	})

	t.Run("No user found", func(t *testing.T) {
		email := "test@test.com"

		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("FindOne", mock.Anything).Return(nil, nil)

		expectedBody, err := json.Marshal(gin.H{
			"user": nil,
		})
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		req, err := http.NewRequest(http.MethodGet, "http://localhost:"+os.Getenv("PORT")+"/api/?email="+email, nil)
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())

	})

	t.Run("No email query", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("FindOne", mock.Anything).Return(nil, nil)

		expectedBody, err := json.Marshal(gin.H{
			"message": "missing email parameter",
		})
		assert.NoError(t, err)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		req, err := http.NewRequest(http.MethodGet, "http://localhost:"+os.Getenv("PORT")+"/api/", nil)
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})
}
