package controllers 

import (
    "net/http"
	"github.com/gin-gonic/gin"
	
	"log"
	"quarxlab/database"
	"quarxlab/models"
	"strconv"
	xerrors "quarxlab/common/errors"
)

func init() {
	database.Database().AutoMigrate(&models.Comment{})
}

type commentController int
const CommentController = commentController(0)

func (this commentController) List(c *gin.Context) {
	
	articleId := c.Param("article_id")
	var article models.Article
	database.Database().First(&article, articleId)

	var comments []models.Comment
	database.Database().Model(&article).Related(&comments)

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": comments})
}

func (this commentController) Create(c *gin.Context) {

	articleId := c.Param("article_id")
	d, _ := strconv.ParseUint(articleId, 0, 64)

	var comment models.Comment
	if err := c.ShouldBind(&comment); err == nil {
		comment.ArticleID = uint(d)
		created := database.Database().Create(&comment).RowsAffected > 0
		if !created {
			err1 := xerrors.NewError(2002)
			panic(&err1)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": nil})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this commentController) Query(c *gin.Context) {

	commentId := c.Param("comment_id")

	var comment models.Comment
	database.Database().First(&comment, commentId)
	if comment.ID == 0 {
		err := xerrors.NewError(2001)
		panic(&err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": comment})
}

func (this commentController) Update(c *gin.Context) {

	commentId := c.Param("comment_id")

	var comment models.Comment

	if err := c.ShouldBind(&comment); err == nil {
		updated := database.Database().Model(&comment).Where("id = ?", commentId).Updates(comment).RowsAffected > 0
		if !updated {
			err := xerrors.NewError(2001)
			panic(&err)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": nil})
		return 
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this commentController) Delete(c *gin.Context) {

	commentId := c.Param("comment_id")
	deleted := database.Database().Where("id = ?", commentId).Delete(&models.Comment{}).RowsAffected > 0
	if !deleted {
		err := xerrors.NewError(2001)
		panic(&err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": nil})
}
