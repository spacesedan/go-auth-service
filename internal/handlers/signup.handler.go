package handlers

import (
	"authentication/internal/apperrors"
	"authentication/internal/dto"
	"authentication/internal/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var rb models.SignUpReq

	if err := c.BindJSON(&rb); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]apperrors.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = apperrors.ErrorMsg{fe.Field(), apperrors.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": out,
			})
		}
		return
	}

	createDto := dto.CreateUser{
		Password: &rb.Password,
		Email:    &rb.Email,
	}

	user, err := h.userService.CreateUser(createDto)
	fmt.Println(err)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.tokenService.GenerateJWT(user)
	if err == apperrors.GeneratingTokenErr {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"user": nil,
		})
		return
	}

	c.SetCookie("user", token, 1000*60*60*24, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})

}
