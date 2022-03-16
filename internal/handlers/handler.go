package handlers

import (
	"authentication/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService service.UserService
}

type Config struct {
	UserService service.UserService

	R *gin.Engine
}

// NewHandler handles the app endpoints
func NewHandler(c *Config) {
	// Place services in here
	h := &Handler{
		userService: c.UserService,
	}

	c.R.GET("/", h.Health)

	api := c.R.Group("api")
	{
		api.GET("/", h.FindUser)
		api.POST("/", h.SignUp)
	}

}
