package controllers

import (
    "net/http"
	"github.com/gin-gonic/gin"
)

type assetsController int
const AssetsController assetsController = 0

func (this assetsController) List(c *gin.Context) {
	c.String(http.StatusOK, "assets list")
}

func (this assetsController) Upload(c *gin.Context) {
	c.String(http.StatusOK, "assets upload")
}

func (this assetsController) Download(c *gin.Context) {
	c.String(http.StatusOK, "assets download")
}

func (this assetsController) Update(c *gin.Context) {
	c.String(http.StatusOK, "assets update")
}

func (this assetsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "assets delete")
}