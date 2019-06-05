package controllers 

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type commentController int
const CommentController commentController = 0

func (this commentController) List(c *gin.Context) {

	c.String(http.StatusOK, "comment list")
}

func (this commentController) Create(c *gin.Context) {

	c.String(http.StatusOK, "comment create")
}

func (this commentController) Query(c *gin.Context) {

    c.String(http.StatusOK, "comment query")
}

func (this commentController) Update(c *gin.Context) {

	c.String(http.StatusOK, "comment update")
}

func (this commentController) Delete(c *gin.Context) {

	c.String(http.StatusOK, "comment delete")
}
