package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nade-harlow/WeekEightTask/Village-square/controllers/helper"
	"github.com/nade-harlow/WeekEightTask/Village-square/models"
	"github.com/nade-harlow/WeekEightTask/Village-square/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var Db *sql.DB
var user models.User

func init() {
	Db, _ = database.MySqlCon()
}

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func SignUpForm(c *gin.Context) {

	fname := c.PostForm("fname")
	lname := c.PostForm("lname")
	email := c.PostForm("email")
	pword := c.PostForm("pword")
	rpword := c.PostForm("rpword")
	helper.VerifyEmpty(fname, lname, c)
	//if fname == "" {
	//	c.String(400, "Invalid first name")
	//	return
	//}
	//if lname == "" {
	//	c.String(400, "Invalid last name")
	//	return
	//}

	VerfiedEmail := helper.IsEmailValid(email, c)

	var v models.User
	Db.QueryRow("SELECT FROM users WHERE email = ?", VerfiedEmail).Scan(&v.UserId, v.FirstName, v.LastName, v.Email, v.Password, v.TimeCreated)
	if VerfiedEmail == v.Email {
		c.String(400, "Email already exist")
		return
	} ////models.DbModel.SignUpHandler(user)
	if len(pword) < 6 || len(rpword) < 6 {
		c.String(400, "password must contain at least six characters!")
		return
		//c.String(400, "password must contain at least six characters!\n")
	}
	if pword != rpword {
		c.String(400, "password mismatched\n")
		return
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pword), bcrypt.MinCost)
		users := &models.User{
			UserId:      uuid.New().String(),
			FirstName:   fname,
			LastName:    lname,
			Email:       email,
			Password:    string(hashedPassword),
			TimeCreated: time.Now().Format(time.RFC850),
		}

		stmt, err := Db.Prepare("INSERT INTO users (id, first_name, last_name, email, password, time_created) VALUE (?,?,?,?,?,?)")
		defer stmt.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = stmt.Exec(users.UserId, users.FirstName, users.LastName, users.Email, users.Password, users.TimeCreated)
		//log.Println(user.TimeCreated)
		if err != nil {
			log.Println(err.Error())
			c.String(400, "unable to insert data")
			return
		}
	}
	c.String(200, "signup successfully")

}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginForm(c *gin.Context) {
	email := c.PostForm("email")
	pword := c.PostForm("pword")

	//dbM := models.DbModel{}
	//dbM.LoginHandler(user.Email, user.Password, c)

	Db.QueryRow("select * from users where email= ?", email).Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.TimeCreated)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pword))
	log.Println(user.Password, pword)
	if er != nil {
		log.Println("wrong username and password: ", er)
		c.String(406, "wrong username and password")
		return
	}

	c.String(200, "Login successful!")
}

func AddPost(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := Db.Prepare(fmt.Sprintf("INSERT INTO posts(id, title, boby, time_created, user_id) VALUES(?, ?, ?, ?)"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	post.Id = "ad01eefe-aaad-45b1-b5ce-2a2b1e53b92c"
	post.TimeCreated = time.Now().Format(time.RFC850)

	res, err := stmt.Exec(post.Id, post.Title, post.Body, post.TimeCreated)
	if err != nil {
		log.Println(err.Error())
		return
	}
	total, _ := res.RowsAffected()
	if total < 1 {
		log.Println(err.Error())
		return
	}
	c.JSON(200, gin.H{"message": "Okay"})
	//if post.Title == "" {
	//	c.JSON(400, log.Println(err.Error())
	//		returngin.H{"message":"enter title fields"})
	//	//c.Redirect(302, "...")
	//}
	//if post.Body == "" {
	//	//http.Error(w, "Enter Content field", 301)
	//	c.JSON(400, gin.H{"message":"enter body field"})
	//	//c.Redirect(302, "...")
	//} else {
	//	log.Println(post.Body,post.Title)
	//	add := models.Post{
	//		Id:      uuid.New().String(),
	//		Title:   post.Title,
	//		Body: 	 post.Body,
	//		TimeCreated:    time.Now().Format(time.RFC850),
	//	}
	//	//
	//	//statement:= `INSERT INTO posts (id, title, body, time_created) VALUE (?,?,?,?)`
	//	//
	//
	//	stmt, er := Db.Prepare("INSERT INTO posts (id, title, body, time_created) VALUES (?,?,?,?)")
	//	if er != nil{
	//		log.Println(err)
	//		c.JSON(500, gin.H{"message": "can't insert post to database "})
	//		return
	//	}
	//
	//	//defer stmt.Close()
	//
	//	fmt.Println("copied")
	//	_, errr := stmt.Exec(add.Id, add.Title, add.Body, add.TimeCreated)
	//	fmt.Println("i got here")
	//
	//	if errr != nil {
	//		log.Println("unable to insert data ", er)
	//		c.JSON(500, gin.H{"message":"unable to insert data"})
	//
	//	}else {
	//
	//		c.JSON(200, gin.H{"message":"post uploaded successfully"})
	//	}

	//c.Redirect(302,"")

}