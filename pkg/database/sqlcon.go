package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func MySqlCon() (db *sql.DB, err error) {

	//host := os.Getenv("HOST")
	//port := os.Getenv("DB_PORT")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	//
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db, nil
}
