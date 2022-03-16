package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignOut(c *gin.Context) {
	userToken, err := c.Cookie("user")
	if err == http.ErrNoCookie {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{})
	}
	c.SetCookie("user", userToken, -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{})
}
