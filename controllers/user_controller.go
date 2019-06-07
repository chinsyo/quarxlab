package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strconv"
	xjwt "quarxlab/common/jwt"
	xerrors "quarxlab/common/errors"
	"golang.org/x/crypto/bcrypt"
	"quarxlab/database"
	"quarxlab/models"
)

func init() {
	database.Database().AutoMigrate(&models.User{}, &models.Profile{})
}

type userController int 
const UserController = userController(0)

func (this userController) check(username, password string) {
	if username == "" || password == "" {
		err := xerrors.NewError(4101)
		panic(&err)
	}
}

func (this userController) Signup(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	this.check(username, password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var user models.User
	if err := c.ShouldBind(&user); err == nil {
		user.Password = string(hash)

		tx := database.Database().Begin()

		database.Database().Create(&user)
		created := database.Database().NewRecord(user)
		if !created {
			tx.Rollback()
			errJson := xerrors.NewError(4002)
			panic(&errJson)
			return
		}

		profile := models.Profile{UserID: user.ID}
		database.Database().Create(&profile)
		created = database.Database().NewRecord(profile)
		if !created {
			tx.Rollback()
			errJson := xerrors.NewError(4002)
			panic(&errJson)
			return
		}
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this userController) Signin(c *gin.Context) {

	
	username := c.PostForm("username")
	password := c.PostForm("password")
	this.check(username, password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	user := models.User{Username: username}
	database.Database().First(&user)
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	userID := strconv.Itoa(int(user.ID))
	jwt, _ := xjwt.GenerateToken(userID)
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": gin.H{"token": jwt}})
}

func (this userController) Logout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": gin.H{}})
}

func (this userController) Verify(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": gin.H{}})
}

func (this userController) Forgot(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	this.check(username, password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var user models.User
	if err := c.ShouldBind(&user); err == nil {
		user.Password = string(hash)
		created := database.Database().Create(&user).RowsAffected > 0
		if !created {
			errJson := xerrors.NewError(4002)
			panic(&errJson)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
	} else {
		log.Fatal(err)
		panic(err)
	}
}

func (this userController) Profile(c *gin.Context) {
	
	if uid := c.Param("user_id"); uid != "" {
		
		var user models.User
		database.Database().Model(&user).Where("id = ?", uid)
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": user.Profile})
		return
	}

	token := c.Query("token")
	claims, _ := xjwt.ParseToken(token)
	userID := claims.UserID

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": gin.H{"uid": userID, }})
}

func (this userController) Edit(c *gin.Context) {

	token := c.Query("token")
	claims, _ := xjwt.ParseToken(token)
	userID, _ := strconv.ParseUint(claims.UserID, 0, 64)

	var user models.User
	var oldProfile models.Profile 
	var newProfile = models.Profile{UserID: uint(userID)}
	if err := c.ShouldBind(&newProfile); err == nil {
		database.Database().Model(&user).Where("user_id = ?", userID).Related(&oldProfile)
		updated := database.Database().Model(&user).Where("user_id = ?", userID).Related(&oldProfile).Updates(&newProfile).RowsAffected > 0

		if !updated {
			errJson := xerrors.NewError(4001)
			panic(&errJson)
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "data": gin.H{}})
		return 
	} else {
		log.Fatal(err)
		panic(err)
	}
}