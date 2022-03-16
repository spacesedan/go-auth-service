package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CurrentUser(c *gin.Context) {
	token, _ := c.Cookie("user")

	user, _ := h.tokenService.ParseJWT(token)
	fmt.Println(user)

	c.JSON(http.StatusOK, gin.H{
		"currentUser": user,
	})

}
