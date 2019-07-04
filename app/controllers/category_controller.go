package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
	"quarxlab/app/database"
	"quarxlab/app/models"
	xerrors "quarxlab/lib/errors"
)

func init() {
	database.Database().AutoMigrate(&models.Category{})
}

type categoryController int

const CategoryController = categoryController(0)

func (this categoryController) List(c *gin.Context) {

	var categories []models.Category
	database.Database().Find(&categories)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": categories})
}

func (this categoryController) Create(c *gin.Context) {

	var category models.Category
	if err := c.ShouldBind(&category); err == nil {
		created := database.Database().Create(&category).RowsAffected > 0
		if !created {
			errJson := xerrors.NewError(3002)
			panic(errJson)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this categoryController) Query(c *gin.Context) {

	categoryID := c.Param("category_id")

	var category models.Category
	database.Database().First(&category, categoryID)

	var articles []models.Article
	database.Database().Model(&category).Related(&articles)

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": articles})
}

func (this categoryController) Update(c *gin.Context) {

	categoryID := c.Param("category_id")

	var category models.Category
	if err := c.ShouldBind(&category); err == nil {
		updated := database.Database().Model(&category).Where("id = ?", categoryID).Updates(category).RowsAffected > 0
		if !updated {
			errJson := xerrors.NewError(3001)
			panic(errJson)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
		return
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this categoryController) Delete(c *gin.Context) {

	categoryID := c.Param("category_id")

	deleted := database.Database().Where("id = ?", categoryID).Delete(&models.Category{}).RowsAffected > 0
	if !deleted {
		errJson := xerrors.NewError(3001)
		panic(errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}
