package database

import "database/sql"

const (
	dbDriver = "mysql"
	dbUser = "root"
	dbPass = "asdfghjk"
	dbName = "villagesquare"
)


func MySqlCon() (db *sql.DB, err error) {
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil{
		return nil, err
	}
	if err:= db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}

