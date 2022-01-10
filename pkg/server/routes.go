package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/WeekEightTask/Village-square/controllers/handlers"
	"github.com/nade-harlow/WeekEightTask/Village-square/middleware"
)

func Routes(r *gin.Engine) {

	userRouter := r.Group("/")
	//userRouter.Use(middleware.CheckLogin())
	{
		//userRouter.GET("/login", middleware.CheckNotLogedin(), handlers.Login)
		userRouter.GET("/signup", middleware.CheckNotLogedin(), handlers.SignUp)
		userRouter.POST("/form", middleware.CheckNotLogedin(), handlers.SignUpForm)
		userRouter.POST("/login/form", middleware.CheckNotLogedin(), handlers.LoginForm)
	}

	r.GET("/", middleware.CheckNotLogedin(), handlers.Login)
	postRouter := r.Group("/post")
	postRouter.Use(middleware.CheckLogin())
	{
		postRouter.GET("/logout", handlers.Logout)
		postRouter.GET("/home", handlers.GetPost)
		postRouter.GET("/home/:Id", handlers.VeiwPost)
		postRouter.GET("/user", handlers.UserPage)
		postRouter.GET("/delete/:Id", handlers.DeletePost)
		postRouter.GET("/edit/:Id", handlers.EditPost)
		postRouter.POST("/edit/process/:Id", handlers.EditPostProcess)
		postRouter.POST("/comment/:Id", handlers.AddComment)
		postRouter.GET("/createpost", handlers.CreatePost)
		postRouter.POST("/createpost/form", handlers.CreatePostProcess)
	}

}
