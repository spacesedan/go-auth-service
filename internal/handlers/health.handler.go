package handlers

import (
	"github.com/gin-gonic/gin"
)

// Health endpoint that returns the health of the server
func (h *Handler) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
