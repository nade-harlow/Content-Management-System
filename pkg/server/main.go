package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nade-harlow/WeekEightTask/Village-square/controllers/handlers"
	"github.com/nade-harlow/WeekEightTask/Village-square/pkg/database"
	"log"
)

func main() {
	db, err := database.MySqlCon()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("./ui/html/*")

	r.GET("/", handlers.SignUp)
	r.POST("/form", handlers.SignUpForm)
	r.GET("/login", handlers.Login)
	r.POST("/login/form", handlers.LoginForm)
	r.POST("/add", handlers.AddPost)
	er := r.Run()
	if er != nil {
		return
	}
}
