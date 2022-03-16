package handlers

import (
	"authentication/internal/models"
	"authentication/internal/models/mocks"
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

func TestCurrentUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		var mockedUser *models.User
		mockedUser = &models.User{
			Email: email,
		}

		mockTokenService := new(mocks.MockedTokenService)
		mockTokenService.On("ParseJWT", mock.Anything).Return(mockedUser, nil)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:            app,
			TokenService: mockTokenService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"currentUser": mockedUser,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "http://localhost:"+os.Getenv("PORT")+"/api/currentUser", nil)
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		fmt.Println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())

	})

	t.Run("No Current User", func(t *testing.T) {

		mockTokenService := new(mocks.MockedTokenService)
		mockTokenService.On("ParseJWT", mock.Anything).Return(nil, nil)

		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R:            app,
			TokenService: mockTokenService,
		})

		expectedBody, err := json.Marshal(gin.H{
			"currentUser": nil,
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "http://localhost:"+os.Getenv("PORT")+"/api/currentUser", nil)
		assert.NoError(t, err)

		app.ServeHTTP(w, req)

		fmt.Println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())

	})
}
