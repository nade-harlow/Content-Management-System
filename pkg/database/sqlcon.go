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

	db, err = sql.Open("postgres", "postgres://zcxydxjqbdeeqe:4d3c5682ea4de184ab2e59f26b5ff9981196ab729e615da5a8d91056198f5c4c@ec2-63-32-7-190.eu-west-1.compute.amazonaws.com:5432/dch1l4n831j80m")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db, nil
}
