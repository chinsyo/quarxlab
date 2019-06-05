package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strings"
	xerrors "quarxlab/common/errors"
)

type userController int 
const UserController userController = 0

func (this userController) Signup(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		err := &xerrors.ErrUnknown
		panic(err)
	}

	c.String(http.StatusOK, "user register")
}

func (this userController) Signin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		err := &xerrors.ErrUnknown
		panic(err)
	}

	c.String(http.StatusOK, "user login")
}

func (this userController) Logout(c *gin.Context) {

	c.String(http.StatusOK, "user logout")
}

func (this userController) Verify(c *gin.Context) {

	c.String(http.StatusOK, "user verify")
}

func (this userController) Forgot(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		err := &xerrors.ErrUnknown
		panic(err)
	}

	c.String(http.StatusOK, "user forgot")
}

func (this userController) Profile(c *gin.Context) {
	
	if uid := c.Param("user_id"); uid != "" {
		profile := fmt.Sprintf("profile uid:%s", uid)
		c.String(http.StatusOK, profile)
		return
	}

	c.String(http.StatusOK, "self profile")
}

func (this userController) Edit(c *gin.Context) {

	c.String(http.StatusOK, "user edit")
}