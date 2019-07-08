package main

import (
	"github.com/gin-gonic/gin"
	. "quarxlab/app/controllers"
	xmiddlewares "quarxlab/lib/middlewares"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(xmiddlewares.Recovery())
	// make sure that dummy is before jwt
	router.Use(xmiddlewares.Dummy())
	router.Use(xmiddlewares.JWT())

	v1 := router.Group("/api/v1")
	{
		// version
		v1.GET("/version", VersionController.Latest)

		// captcha
		v1.GET("/captcha", CaptchaController.Refresh)
		v1.POST("/captcha", CaptchaController.Verify)

		// assets
		v1.GET("/assets", AssetsController.List)
		v1.POST("/assets", AssetsController.Upload)
		v1.DELETE("/assets/:assets_id", AssetsController.Delete)

		// users
		v1.POST("/signup", UserController.Signup)
		v1.POST("/signin", UserController.Signin)
		v1.POST("/logout", UserController.Logout)
		v1.POST("/verify", UserController.Verify)
		v1.POST("/forgot", UserController.Forgot)
		v1.POST("/profile", UserController.Edit)
		v1.GET("/profile", UserController.Profile)
		v1.GET("/profile/:user_id", UserController.Profile)

		// articles
		v1.GET("/article", ArticleController.List)
		v1.POST("/article", ArticleController.Create)
		v1.GET("/article/:article_id", ArticleController.Query)
		v1.PUT("/article/:article_id", ArticleController.Update)
		v1.DELETE("/article/:article_id", ArticleController.Delete)

		// comments
		v1.GET("/article/:article_id/comment", CommentController.List)
		v1.POST("/article/:article_id/comment", CommentController.Create)
		v1.GET("/comment/:comment_id", CommentController.Query)
		v1.PUT("/comment/:comment_id", CommentController.Update)
		v1.DELETE("/comment/:comment_id", CommentController.Delete)

		// categories
		v1.GET("/category", CategoryController.List)
		v1.POST("/category", CategoryController.Create)
		v1.GET("/category/:category_id", CategoryController.Query)
		v1.PUT("/category/:category_id", CategoryController.Update)
		v1.DELETE("/category/:category_id", CategoryController.Delete)
	}
	router.Run(":8000")
}
