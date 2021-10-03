package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	UserId string `json:"user_id"`
	FirstName string 	`json:"first_name"`
	LastName string 	`json:"last_name"`
	Email string
	Password string `json:"password"`
	ConfrimPassword string `json:"confrim_password"`
	TimeCreated  string `json:"time_created"`
}

type DbModel struct {
	Db *sql.DB
}

func (db *DbModel) SignUpHandler(u *User)  {

	stmt, err := db.Db.Prepare("INSERT INTO users (id, first_name, last_name, email, password, time_created) VALUE (?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal("unable to insert data ", err)
	}
	_, err = stmt.Exec(u.UserId, u.FirstName, u.LastName, u.Email, u.Password, u.TimeCreated)
	if err != nil {
		log.Fatal("unable to insert data ", err)
	}

}

func (db *DbModel) LoginHandler(email ,password string)  {
	//row:= db.Db.QueryRow("select * from users where first_name= ?", email )
	//temp:= &User{}
	//err :=row.Scan(&temp.Password)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//err = bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(password))
	//if err != nil{
	//	log.Fatal("wrong username and password", err)
	//}

	temp:= &User{}
	db.Db.QueryRow("select * from users where first_name= ?",email ).Scan(&temp.UserId,&temp.FirstName,&temp.LastName,&temp.Email,&temp.Password,&temp.TimeCreated)

	err := bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(password))
	if err != nil{
		log.Fatal("wrong username and password", err)
	}

}