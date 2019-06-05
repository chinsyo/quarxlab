package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"strings"
	xerrors "quarxlab/common/errors"
	"golang.org/x/crypto/bcrypt"
	"quarxlab/database"
	"quarxlab/models"
)

func init() {
	database.Database().AutoMigrate(&models.User{}, &models.Credential{})
}

type userController int 
const UserController = userController(0)

func (this userController) Signup(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		err := &xerrors.ErrUnknown
		panic(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("password", string(hash))
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}

func (this userController) Signin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		err := &xerrors.ErrUnknown
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}

func (this userController) Logout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}

func (this userController) Verify(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}

func (this userController) Forgot(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		err := &xerrors.ErrUnknown
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}

func (this userController) Profile(c *gin.Context) {
	
	if uid := c.Param("user_id"); uid != "" {
		profile := fmt.Sprintf("profile uid:%s", uid)
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{uid: profile}})

		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}

func (this userController) Edit(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": map[string]interface{}{}})
}