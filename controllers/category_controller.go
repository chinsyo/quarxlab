package controllers 

import (
    "net/http"
	"github.com/gin-gonic/gin"
	
	"quarxlab/database"
	"quarxlab/models"
	"log"
	xerrors "quarxlab/common/errors"
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
			panic(&errJson)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this categoryController) Query(c *gin.Context) {

	categoryId := c.Param("category_id")

	var category models.Category
    database.Database().First(&category, categoryId)
    
    var articles []models.Article
    database.Database().Model(&category).Related(&articles)

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": articles})
}

func (this categoryController) Update(c *gin.Context) {

	categoryId := c.Param("category_id")

	var category models.Category
	if err := c.ShouldBind(&category); err == nil {
		updated := database.Database().Model(&category).Where("id = ?", categoryId).Updates(category).RowsAffected > 0
		if !updated {
			err := xerrors.NewError(3001)
			panic(&err)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
		return 
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this categoryController) Delete(c *gin.Context) {

	categoryId := c.Param("category_id")

	deleted := database.Database().Where("id = ?", categoryId).Delete(&models.Category{}).RowsAffected > 0
	if !deleted {
		err := xerrors.NewError(3001)
		panic(&err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}
