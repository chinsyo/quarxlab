package controllers

import (
	"io"
	"os"
	"log"
	"fmt"
    "net/http"
	"github.com/gin-gonic/gin"
	"github.com/bwmarrin/snowflake"
	"quarxlab/app/database"
	"quarxlab/app/models"
)

func init() {
	database.Database().AutoMigrate(&models.Asset{})
}

type assetsController int
const AssetsController = assetsController(0)

func (this assetsController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}

func (this assetsController) Upload(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	filename := fmt.Sprintf("%d_%s", node.Generate().Int64(), header.Filename)

	dir, err := os.Create("static/upload/" + filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer dir.Close()

	_, err = io.Copy(dir, file)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}

func (this assetsController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}