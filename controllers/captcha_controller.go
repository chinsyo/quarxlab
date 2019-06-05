package controllers

import (
    "github.com/gin-gonic/gin"
    // "github.com/dchest/captcha"
    "net/http"
)

type captchaController int
const CaptchaController = captchaController(0)

func (this captchaController) Refresh(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"code": 0})
}

func (this captchaController) Verify(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"code": 0})
}
