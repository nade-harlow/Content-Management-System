package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	Routes(r)
	r.LoadHTMLGlob("./ui/html/*")
	er := r.Run()
	if er != nil {
		return
	}
}
