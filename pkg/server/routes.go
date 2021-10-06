package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/WeekEightTask/Village-square/controllers/handlers"
	"github.com/nade-harlow/WeekEightTask/Village-square/middleware"
)

func Routes(r *gin.Engine) {

	userRouter := r.Group("/")
	userRouter.Use(middleware.CheckLogin())
	{
		userRouter.GET("/login", middleware.CheckNotLogedin(), handlers.Login)
		userRouter.GET("/signup", middleware.CheckNotLogedin(), handlers.SignUp)
		userRouter.POST("/form", middleware.CheckNotLogedin(), handlers.SignUpForm)
		userRouter.POST("/login/form", middleware.CheckNotLogedin(), handlers.LoginForm)
	}

	loog := r.Group("/")
	loog.Use(middleware.CheckNotLogedin())
	{
		//loog.GET("/logout", handlers.Logout)
		//loog.GET("/login", middleware.CheckNotLogedin(),handlers.Login)
	}

	postRouter := r.Group("/post")
	postRouter.Use(middleware.CheckLogin())
	{
		postRouter.GET("/logout", handlers.Logout)
		postRouter.GET("/home", handlers.GetPost)
		postRouter.GET("/pro", handlers.User)
		postRouter.GET("/createpost", handlers.CreatePost)
		postRouter.POST("/createpost/form", handlers.CreatePostProcess)
	}

}
