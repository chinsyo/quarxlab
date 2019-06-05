package controllers

import (
	"io"
	"os"
	"log"
	"fmt"
	"time"
    "net/http"
	"github.com/gin-gonic/gin"
	"quarxlab/database"
	"quarxlab/models"
)

func init() {
	database.Database().AutoMigrate(&models.Asset{})
}

type assetsController int
const AssetsController = assetsController(0)

func (this assetsController) List(c *gin.Context) {
	c.String(http.StatusOK, "assets list")
}

func (this assetsController) Upload(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusBadRequest, "Bad request")
	}

	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("%d_%s", timestamp, header.Filename)

	dir, err := os.Create("static/upload/" + filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer dir.Close()

	_, err = io.Copy(dir, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful")
}

func (this assetsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "assets delete")
}