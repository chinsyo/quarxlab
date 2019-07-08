package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dummy() gin.HandlerFunc {
	return func(c *gin.Context) {

		statusCode := c.Writer.Status()
		if statusCode != http.StatusOK {
			//c.Next()
			c.Abort()
			return
		}
		c.Next()
	}
}
