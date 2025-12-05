package middlewares

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func RequestFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains([]string{"GET", "POST", "PUT", "DELETE"}, c.Request.Method) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	}
}
