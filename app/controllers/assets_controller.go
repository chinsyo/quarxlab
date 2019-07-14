package controllers

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"quarxlab/app/database"
	"quarxlab/app/models"
	xerrors "quarxlab/lib/errors"
)

func init() {
	database.Database().AutoMigrate(&models.Asset{})
}

type assetsController int

const AssetsController = assetsController(0)

func (this assetsController) List(c *gin.Context) {
	var assets []models.Asset
	database.Database().Find(&assets)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": assets})
}

func (this assetsController) Submit(c *gin.Context) {

	// filePath := c.Param("file_path")
	// log.Printf("filePath", filePath)

	var asset models.Asset
	if err := c.ShouldBind(&asset); err == nil {

		updated := database.Database().Save(&asset).RowsAffected > 0
		if !updated {
			errJson := xerrors.NewError(1001)
			panic(errJson)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": asset})
		return
	} else {
		log.Fatal(err)
		panic(err)
	}

	// c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}

func (this assetsController) Upload(c *gin.Context) {

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	filename := fmt.Sprintf("%d_%s", node.Generate().Int64(), handler.Filename)
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

	asset := models.Asset{FilePath: "static/upload/" + filename}
	created := database.Database().Create(&asset).RowsAffected > 0
	if !created {
		errJson := xerrors.NewError(1002)
		panic(errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{"id": asset.ID}})
}

func (this assetsController) Delete(c *gin.Context) { 
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": ""})
}
