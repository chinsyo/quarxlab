package middlewares

import (
    "github.com/gin-gonic/gin"
    xerrors "quarxlab/common/errors"
    "net/http"
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
                        c.JSON(http.StatusInternalServerError, &xerrors.ErrUnknown)
                }
            }
        }()
        c.Next()
    }
}