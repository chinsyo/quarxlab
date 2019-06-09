package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strconv"
	xjwt "quarxlab/lib/jwt"
	xerrors "quarxlab/lib/errors"
	xctx "quarxlab/lib/context"
	"golang.org/x/crypto/bcrypt"
	"quarxlab/app/database"
	"quarxlab/app/models"
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
		created := database.Database().Create(&user).RowsAffected > 0
		if !created {
			tx.Rollback()
			log.Fatal("User create failed")
			errJson := xerrors.NewError(4002)
			panic(&errJson)
		}

		profile := models.Profile{UserID: user.ID}
		created = database.Database().Create(&profile).RowsAffected > 0
		if !created {
			tx.Rollback()
			log.Fatal("Profile create failed")
			errJson := xerrors.NewError(4002)
			panic(&errJson)
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

	jwt, _ := xjwt.GenerateToken(user.ID)
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
	
	var userID uint
	if uid := c.Param("user_id"); uid != "" {
		tmpUID, _ := strconv.ParseUint(uid, 0, 0)
		userID = uint(tmpUID)
	} else {
		ctxUID, _ := c.Get(xctx.UID)
		userID = ctxUID.(uint)
	}
	var profile = models.Profile{UserID: userID}
	database.Database().First(&profile)
	if profile.ID == 0 {
		errJson := xerrors.NewError(4001)
		panic(&errJson)
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": nil, "data": profile })
}

func (this userController) Edit(c *gin.Context) {

	userID, _ := c.Get(xctx.UID)

	var oldProfile models.Profile 
	var newProfile = models.Profile{UserID: userID.(uint)}
	if err := c.ShouldBind(&newProfile); err == nil {
		updated := database.Database().Model(&oldProfile).Where("user_id = ?", userID).Updates(&newProfile).RowsAffected > 0
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