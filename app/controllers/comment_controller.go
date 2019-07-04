package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
	"quarxlab/app/database"
	"quarxlab/app/models"
	xerrors "quarxlab/lib/errors"
	"strconv"
)

func init() {
	database.Database().AutoMigrate(&models.Comment{})
}

type commentController int

const CommentController = commentController(0)

func (this commentController) List(c *gin.Context) {

	articleID := c.Param("article_id")
	var article models.Article
	database.Database().First(&article, articleID)

	var comments []models.Comment
	database.Database().Model(&article).Related(&comments)

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": comments})
}

func (this commentController) Create(c *gin.Context) {

	articleID := c.Param("article_id")
	d, _ := strconv.ParseUint(articleID, 0, 64)

	var comment models.Comment
	if err := c.ShouldBind(&comment); err == nil {
		comment.ArticleID = uint(d)
		created := database.Database().Create(&comment).RowsAffected > 0
		if !created {
			errJson := xerrors.NewError(2002)
			panic(errJson)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this commentController) Query(c *gin.Context) {

	commentID := c.Param("comment_id")

	var comment models.Comment
	database.Database().First(&comment, commentID)
	if comment.ID == 0 {
		errJson := xerrors.NewError(2001)
		panic(errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": comment})
}

func (this commentController) Update(c *gin.Context) {

	commentID := c.Param("comment_id")

	var comment models.Comment

	if err := c.ShouldBind(&comment); err == nil {
		updated := database.Database().Model(&comment).Where("id = ?", commentID).Updates(comment).RowsAffected > 0
		if !updated {
			errJson := xerrors.NewError(2001)
			panic(errJson)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
		return
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this commentController) Delete(c *gin.Context) {

	commentID := c.Param("comment_id")
	deleted := database.Database().Where("id = ?", commentID).Delete(&models.Comment{}).RowsAffected > 0
	if !deleted {
		errJson := xerrors.NewError(2001)
		panic(errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
}
