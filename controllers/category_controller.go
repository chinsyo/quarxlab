package controllers 

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type categoryController int
const CategoryController categoryController = 0

func (this categoryController) List(c *gin.Context) {

	c.String(http.StatusOK, "category list")
}

func (this categoryController) Create(c *gin.Context) {

	c.String(http.StatusOK, "category insert")
}

func (this categoryController) Query(c *gin.Context) {

    c.String(http.StatusOK, "category query")
}

func (this categoryController) Update(c *gin.Context) {

	c.String(http.StatusOK, "category update")
}

func (this categoryController) Delete(c *gin.Context) {

	c.String(http.StatusOK, "category delete")
}
