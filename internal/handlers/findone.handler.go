package handlers

import (
	"authentication/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) FindUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing email parameter",
		})
		return
	}

	findOne := dto.FindOne{
		Email: &email,
	}

	user, err := h.userService.FindOne(findOne)
	if err != nil || user == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}
