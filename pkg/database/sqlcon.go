package database

import (
	"database/sql"
	"fmt"
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

	db, err = sql.Open("postgres", "postgres://etfnhypcroxzcr:e1e63cc4807e17d065ce407470611f7861a32be49fd52754d3428ec9c5cb09b9@ec2-18-203-64-130.eu-west-1.compute.amazonaws.com:5432/d453tbtrt69gc8")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db, nil
}
