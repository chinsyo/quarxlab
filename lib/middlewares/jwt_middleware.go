package middlewares

import (
    "github.com/gin-gonic/gin"
    xjwt "quarxlab/lib/jwt"
    xerrors "quarxlab/lib/errors"
    xctx "quarxlab/lib/context"
    "net/http"
    "time"
)

var whitelist = []string{
    "/api/v1/version",
    "/api/v1/signup",
    "/api/v1/signin",
    "/api/v1/logout",
    "/api/v1/forgot",
    "/api/v1/verify",
}

func bypass(path string) bool {
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

        if bypass(path) {
            c.Next()
            return
        }

        token := c.Query("token")
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
            c.Set(xctx.UID, claims.UserID)
        }
        c.Next()
    }
}