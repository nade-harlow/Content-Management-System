package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/nade-harlow/WeekEightTask/Village-square/pkg/database"
	"github.com/nade-harlow/WeekEightTask/Village-square/pkg/server"
	"log"
	"os"
)

func main() {
	er := godotenv.Load()
	if er != nil {
		log.Println(er.Error())
	}
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port += "8080"
	}
	db, err := database.MySqlCon()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	server.Routes(r)
	r.LoadHTMLGlob("./ui/html/*")
	er = r.Run(port)
	if er != nil {
		return
	}
}
