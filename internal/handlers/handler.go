package handlers

import (
	"authentication/internal/middleware"
	"authentication/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService  service.UserService
	tokenService service.TokenService
}

type Config struct {
	UserService  service.UserService
	TokenService service.TokenService

	R *gin.Engine
}

// NewHandler handles the app endpoints
func NewHandler(c *Config) {
	// Place services in here
	h := &Handler{
		userService:  c.UserService,
		tokenService: c.TokenService,
	}

	c.R.GET("/", h.Health)

	api := c.R.Group("api")
	{
		if gin.Mode() != gin.TestMode {
			api.GET("/", h.FindUser)
			api.POST("/", h.SignUp)
			api.POST("/signin", h.Signin)
			api.GET("/signout", h.SignOut)
			api.GET("/currentUser", middleware.CurrentUser(), h.CurrentUser)
		} else {
			api.GET("/", h.FindUser)
			api.POST("/", h.SignUp)
			api.POST("/signin", h.Signin)
			api.GET("/signout", h.SignOut)
			api.GET("/currentUser", h.CurrentUser)
		}

	}

}
