package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()

		app := gin.Default()

		NewHandler(&Config{
			R: app,
		})

		expectedBody, err := json.Marshal(gin.H{
			"status": "ok",
		})
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.NoError(t, err)

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.Bytes())

	})
}
