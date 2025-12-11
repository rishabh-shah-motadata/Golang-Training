package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodPost || c.Request.Method == http.MethodDelete || c.Request.Method == http.MethodPut {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	}
}
