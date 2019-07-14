package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		statusCode := c.Writer.Status()
		if statusCode != http.StatusOK {
			c.Abort()
			return
		}
		c.Next()
	}
}
