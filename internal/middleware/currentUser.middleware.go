package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("user")
		if err == http.ErrNoCookie {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"currentUser": nil,
			})
			return
		}
		c.Next()

	}
}
