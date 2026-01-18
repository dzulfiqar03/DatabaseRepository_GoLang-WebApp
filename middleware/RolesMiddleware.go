package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RolesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("Role")

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}