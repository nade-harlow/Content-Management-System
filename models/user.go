package models

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/nade-harlow/WeekEightTask/Village-square/controllers/handlers"
	"log"
)

type User struct {
	UserId          string `json:"user_id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfrimPassword string `json:"confrim_password"`
	TimeCreated     string `json:"time_created"`
}

type DbModel struct {
	Db *sql.DB
}
type UserCRUD interface {
	CreateUsers()
}

func (db *DbModel) SignUpHandler(u *User) {

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

func (db *DbModel) LoginHandler(email, password string, c *gin.Context) {
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

	//var user User
	//dummy:= user.Password
	//db.Db.QueryRow("select * from users where email= ?",user.Email ).Scan(&user.UserId,&user.FirstName,&user.LastName,&user.Email,&user.Password,&user.TimeCreated)
	//
	//er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dummy))
	//log.Println(user.Password, dummy)
	//if er != nil{
	//	log.Println("wrong username and password: ", er)
	//	c.String(406, "wrong username and password")
	//	return
	//}
	//c.String(200,"Login successful!")

}

func CreateUsers(ctx *gin.Context, a, b, c, d, e, f string) {

	stmt, err := handlers.Db.Prepare("INSERT INTO users (id, first_name, last_name, email, password, time_created) VALUE (?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, err = stmt.Exec(a, b, c, d, e, f)
	//log.Println(user.TimeCreated)
	if err != nil {
		log.Println(err.Error())
		ctx.String(400, "unable to insert data")
		return
	}
}
