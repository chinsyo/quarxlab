package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	xctx "quarxlab/lib/context"
	xerrors "quarxlab/lib/errors"
	xjwt "quarxlab/lib/jwt"
	"time"
	"log"
)

var whitelist = []string{
	"/api/v1/version",
	"/api/v1/signup",
	"/api/v1/signin",
	"/api/v1/logout",
	"/api/v1/forgot",
	"/api/v1/verify",
	"/api/v1/captcha",
	"/api/v1/assets",
}

func shouldBypass(path string) bool {
	for _, v := range whitelist {
		if v == path {
			return true
		}
	}
	return false
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		if shouldBypass(path) {
			c.Next()
			return
		}

		// token := c.Query("token")
		token := c.GetHeader("X-Auth-Token")
		errJson := xerrors.ErrSuccess

		if token == "" {
			errJson = xerrors.ErrTokenRequired
		}

		claims, err := xjwt.ParseToken(token)
		if err != nil {
			errJson = xerrors.ErrTokenFailure
		} else if time.Now().Unix() > claims.ExpiresAt {
			errJson = xerrors.ErrTokenExpired
		}

		if errJson != xerrors.ErrSuccess {
			c.JSON(http.StatusUnauthorized, gin.H{"code": errJson.Code, "message": errJson.Message, "data": nil})
			c.Abort()
			return
		} else {
			log.Println("c.Set", claims.UserID)
			c.Set(xctx.UID, claims.UserID)
		}
		c.Next()
	}
}
