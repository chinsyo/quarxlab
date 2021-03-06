package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"quarxlab/app/database"
	"quarxlab/app/models"
	xerrors "quarxlab/lib/errors"
)

func init() {
	database.Database().AutoMigrate(&models.Article{})
}

type articleController int

const ArticleController = articleController(0)

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
			errJson := xerrors.ErrArticlePubFailed
			panic(errJson)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this articleController) Query(c *gin.Context) {

	articleID := c.Param("article_id")

	var article models.Article
	database.Database().First(&article, articleID)
	if article.ID == 0 {
		errJson := xerrors.ErrArticleNotExist
		panic(errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": article})
}

func (this articleController) Update(c *gin.Context) {

	articleID := c.Param("article_id")

	var article models.Article
	if err := c.ShouldBind(&article); err == nil {
		updated := database.Database().Model(&article).Where("id = ?", articleID).Updates(article).RowsAffected > 0
		if !updated {
			errJson := xerrors.ErrArticleNotExist
			panic(errJson)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
		return
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this articleController) Delete(c *gin.Context) {

	articleID := c.Param("article_id")
	deleted := database.Database().Where("id = ?", articleID).Delete(&models.Article{}).RowsAffected > 0
	if !deleted {
		errJson := xerrors.ErrArticleNotExist
		panic(errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}
