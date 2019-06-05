package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quarxlab/models"
)

type versionController int 
const VersionController versionController = 0

func (this versionController) Latest(c *gin.Context) {
	c.String(http.StatusOK, models.Current.String())
}