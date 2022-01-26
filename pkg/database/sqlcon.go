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

	db, err = sql.Open("postgres", "postgres://vysjvbplqhhnkq:a78c377957dbb1342f8e1932f03645a7ec6a8301397ccc14e4169ea05473bedc@ec2-54-170-163-224.eu-west-1.compute.amazonaws.com:5432/dfi5c1hhmnm8k4")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db, nil
}
