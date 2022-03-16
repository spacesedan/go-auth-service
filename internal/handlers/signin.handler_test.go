package handlers

import (
	"authentication/internal/apperrors"
	"authentication/internal/models"
	"authentication/internal/models/mocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSignIn(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		var mockedUser *models.User
		mockedUser = &models.User{
			Email:    email,
			Password: password,
		}

		mockUserService := new(mocks.MockedUserService)
		mockTokenService := new(mocks.MockedTokenService)
		mockUserService.On("SignIn", mock.Anything).Return(mockedUser, nil)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:            app,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"user": mockedUser,
		})
		assert.NoError(t, err)

		testBody, err := json.Marshal(models.SignInReq{
			Email:    email,
			Password: password,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/signin", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

	t.Run("Wrong Password", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("SignIn", mock.Anything).Return(nil, apperrors.WrongPasswordErr)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"error": apperrors.WrongPasswordErr.Error(),
		})
		assert.NoError(t, err)

		testBody, err := json.Marshal(models.SignInReq{
			Email:    email,
			Password: password,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/signin", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

	t.Run("Missing Request Parameters", func(t *testing.T) {
		mockUserService := new(mocks.MockedUserService)
		mockUserService.On("SignIn", mock.Anything).Return(nil, nil)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:           app,
			UserService: mockUserService,
		})

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
		assert.NoError(t, err)

		testBody, err := json.Marshal(models.SignInReq{})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/signin", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

	t.Run("No JWT Token", func(t *testing.T) {
		var mockedUser *models.User
		mockedUser = &models.User{
			Email:    email,
			Password: password,
		}

		mockUserService := new(mocks.MockedUserService)
		mockTokenService := new(mocks.MockedTokenService)
		mockUserService.On("SignIn", mock.Anything).Return(mockedUser, nil)
		mockTokenService.On("GenerateJWT", mock.Anything).Return(nil, apperrors.GeneratingTokenErr)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:            app,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"user": nil,
		})
		assert.NoError(t, err)

		testBody, err := json.Marshal(models.SignInReq{
			Email:    email,
			Password: password,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+os.Getenv("PORT")+"/api/signin", bytes.NewBuffer(testBody))
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())
	})

}
