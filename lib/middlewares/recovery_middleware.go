package middlewares

import (
    "github.com/gin-gonic/gin"
    xerrors "quarxlab/lib/errors"
    "net/http"
    "log"
)

func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                switch err.(type) {
                    case (*xerrors.Error):
                        errJson := err.(*xerrors.Error)
                        c.JSON(http.StatusOK, gin.H{"code": errJson.Code, "message": errJson.Message, "data": nil})
                    default:
                        errJson := &xerrors.ErrUnknown
                        // errJson.Message = err.String()
                        log.Fatal(err)
                        c.JSON(http.StatusInternalServerError, gin.H{"code": errJson.Code, "message": errJson.Message, "data": err})
                }
            }
        }()
        c.Next()
    }
}