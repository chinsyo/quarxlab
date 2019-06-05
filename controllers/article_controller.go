package controllers

import (
    "net/http"
	"github.com/gin-gonic/gin"
	
	"quarxlab/database"
	"quarxlab/models"
	"log"
	xerrors "quarxlab/common/errors"
)

type articleController int
const ArticleController articleController = 0

func (this articleController) List(c *gin.Context) {
	var articles []models.Article
	database.Database().Find(&articles)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": articles})
}

func (this articleController) Create(c *gin.Context) {

	var article models.Article
	if err := c.ShouldBind(&article); err == nil {
		created := database.Database().Create(&article).RowsAffected > 0
		if !created {
			err1 := xerrors.NewError(1002)
			panic(&err1)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": nil})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this articleController) Query(c *gin.Context) {

	id := c.Param("article_id")

	var article models.Article
	database.Database().First(&article, id)
	if article.ID == 0 {
		err := xerrors.NewError(1001)
		panic(&err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": article})
}

func (this articleController) Update(c *gin.Context) {

	id := c.Param("article_id")

	var article models.Article
	database.Database().First(&article, id)

	if err := c.ShouldBind(&article); err == nil {
		updated := database.Database().Save(&article).RowsAffected > 0
		if !updated {
			err := xerrors.NewError(1001)
			panic(&err)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": nil})
		return 
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this articleController) Delete(c *gin.Context) {
	id := c.Param("article_id")
	deleted := database.Database().Where("id = ?", id).Delete(&models.Article{}).RowsAffected > 0
	if !deleted {
		err := xerrors.NewError(1001)
		panic(&err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": nil})
}
