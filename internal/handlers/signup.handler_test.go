package handlers

import (
	"authentication/internal/apperrors"
	"authentication/internal/models"
	"authentication/internal/models/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	email    = "test@test.com"
	password = "password"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {

		var mockedUser *models.User
		mockedUser = &models.User{
			Email:    email,
			Username: "test",
		}

		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("FindOne", mock.Anything).Return(nil, nil)
		mockUserService.On("CreateUser", mock.Anything).Return(mockedUser, nil)

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

		testBody, err := json.Marshal(models.SignUpReq{
			Email:    email,
			Password: password,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())

	})

	t.Run("User already exists", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("CreateUser", mock.Anything).Return(nil, apperrors.AlreadyExists())

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"error": apperrors.AlreadyExists().Error(),
		})

		testBody, err := json.Marshal(models.SignUpReq{
			Email:    email,
			Password: password,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

	t.Run("Missing Password", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("CreateUser", mock.Anything).Return(nil, apperrors.AlreadyExists())

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		testBody, err := json.Marshal(models.SignUpReq{
			Email: email,
		})
		assert.NoError(t, err)

		expectedBody, err := json.Marshal(gin.H{
			"errors": []apperrors.ErrorMsg{
				{
					Field:   "Password",
					Message: "This field is required",
				},
			},
		})

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

	t.Run("Missing Email", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("CreateUser", mock.Anything).Return(nil, apperrors.AlreadyExists())

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		testBody, err := json.Marshal(models.SignUpReq{
			Password: password,
		})
		assert.NoError(t, err)

		expectedBody, err := json.Marshal(gin.H{
			"errors": []apperrors.ErrorMsg{
				{
					Field:   "Email",
					Message: "This field is required",
				},
			},
		})

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

	t.Run("Missing Body", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("CreateUser", mock.Anything).Return(nil, apperrors.AlreadyExists())

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		testBody, err := json.Marshal(models.SignUpReq{})

		expectedBody, err := json.Marshal(gin.H{
			"errors": []apperrors.ErrorMsg{
				{
					Field:   "Email",
					Message: "This field is required",
				},
				{
					Field:   "Password",
					Message: "This field is required",
				},
			},
		})

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		fmt.Println("Test body", w.Body.String())
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})
}
