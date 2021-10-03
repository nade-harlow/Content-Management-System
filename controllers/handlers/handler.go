package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nade-harlow/WeekEightTask/Village-square/models"
	"github.com/nade-harlow/WeekEightTask/Village-square/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
	"time"
)

var Db *sql.DB
var user models.User

func init() {
	Db, _ = database.MySqlCon()
}

func SignUp(c *gin.Context)  {
	c.HTML(http.StatusOK,"signup.html\n",nil)
}

func isEmailValid(e string, c *gin.Context)  string {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if emailRegex.MatchString(e)== false{
		c.String(400, "Invalid email\n")
		c.Abort()
		//c.AbortWithStatus(400)
	}
	return e
}

func SignUpForm(c *gin.Context)  {
	err := c.BindJSON(&user)
	if err != nil{
		log.Println(err)
		return
	}

	if user.FirstName == "" {
		c.String(400, "Invalid first name\n")
	}
	if user.LastName == "" {
		c.String(400, "Invalid last name\n")
	}

	VerfiedEmail:=isEmailValid(user.Email, c)

	log.Println(user.FirstName,user.LastName,user.Email,user.Password)

	if len(user.Password) < 6 || len(user.ConfrimPassword)<6{
		c.JSON(400, gin.H{"message":"password must contain at least six characters!"})
		c.Abort()
		//c.String(400, "password must contain at least six characters!\n")
	}
	if user.Password == user.ConfrimPassword{
		c.String(400, "password mismatched\n")
	}else {
		hashedPassword, err :=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.MinCost)
		users:= &models.User{
			UserId:    uuid.New().String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     VerfiedEmail,
			Password: string(hashedPassword),
			TimeCreated: time.Now().Format(time.RFC850),
		}


		////models.DbModel.SignUpHandler(user)
		//
		////log.Println(string(hashedPassword))

		stmt, err := Db.Prepare("INSERT INTO users (id, first_name, last_name, email, password, time_created) VALUE (?,?,?,?,?,?)")
		defer stmt.Close()
		if err != nil {
			log.Fatal("unable to insert data ", err)
		}
		_, err = stmt.Exec(users.UserId, users.FirstName, users.LastName, users.Email, users.Password, users.TimeCreated)
		//log.Println(user.TimeCreated)
		if err != nil {
			log.Fatal("unable to insert data ", err)
		}
	}


}

func Login(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}

func LoginForm(c *gin.Context)  {
	email:= c.PostForm("email")
	password:= c.PostForm("pword")

	//dbM := models.DbModel{}
	//dbM.LoginHandler(email, password)

	Db.QueryRow("select * from users where first_name= ?",email ).Scan(&user.UserId,&user.FirstName,&user.LastName,&user.Email,&user.Password,&user.TimeCreated)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		log.Fatal("wrong username and password", err)
		c.Status(http.StatusNotAcceptable)
	}

	c.Status(http.StatusFound)
}