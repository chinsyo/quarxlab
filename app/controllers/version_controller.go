package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quarxlab/app/models"
)

type versionController int 
const VersionController = versionController(0)

func (this versionController) Latest(c *gin.Context) {
	latest := models.LatestVersion.String()
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": gin.H{"version": latest}})
}