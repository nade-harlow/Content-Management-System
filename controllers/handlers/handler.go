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

	//checks empty string
	helper.VerifyEmpty(fname, lname, c)
	// checks if email is valid
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
	c.Redirect(302, "/home")

}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginForm(c *gin.Context) {
	email := c.PostForm("email")
	pword := c.PostForm("pword")

	Db.QueryRow("select * from users where email= ?", email).Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.TimeCreated)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pword))
	if er != nil {
		log.Println("wrong username and password: ", er)
		c.String(406, "wrong username and password")
		return
	}
	c.SetCookie("session", user.UserId, 3600, "/", "localhost", false, true)

	c.Redirect(302, "/post/home")
}

func Logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", false, true)
	c.Redirect(302, "/post/login")
}

func CreatePost(c *gin.Context) {
	c.HTML(200, "createpost.html", nil)
}

func CreatePostProcess(c *gin.Context) {
	id, _ := c.Get("userId")
	newId := id.(string)
	title := c.PostForm("title")
	body := c.PostForm("body")
	helper.VerifyEmptyString(title, body, c)

	stmt, err := Db.Prepare(fmt.Sprintf("INSERT INTO posts(id, title, boby, time_created, user_id) VALUES(?, ?, ?, ?,?)"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	var post = models.Post{
		Id:          uuid.New().String(),
		Title:       title,
		Body:        body,
		TimeCreated: time.Now().Format(time.RFC850),
		UserId:      newId,
	}

	res, err := stmt.Exec(post.Id, post.Title, post.Body, post.TimeCreated, post.UserId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	total, _ := res.RowsAffected()
	if total < 1 {
		log.Println(err.Error())
		return
	}
	c.String(200, "Post added successfully")
	c.Redirect(302, "/post/home")

}

func User(c *gin.Context) {
	uu, err := c.Cookie(user.UserId)
	if err != nil {
		log.Println(err.Error())
	}
	c.HTML(200, "Userpage.html", uu)
}

func GetPost(c *gin.Context) {
	rows, err := Db.Query("select * from posts")
	if err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}
	defer rows.Close()
	var row []models.Post
	for rows.Next() {
		var r models.Post
		err = rows.Scan(&r.Id, &r.Title, &r.Body, &r.TimeCreated, &r.UserId)
		if err != nil {
			log.Println(err.Error())
			c.Status(500)
			return
		}
		row = append(row, r)
	}
	c.HTML(200, "Home.html", row)
}
