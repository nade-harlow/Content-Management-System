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
	"strconv"
	"time"
)

var Db *sql.DB
var user models.User

func init() {
	Db, _ = database.MySqlCon()
}

func SignUp(c *gin.Context) {
	helper.Mail()
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
	c.Redirect(302, "/post/home")

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
	c.Redirect(302, "/login")
}

func CreatePost(c *gin.Context) {
	c.HTML(200, "createpost.html", nil)
}

func CreatePostProcess(c *gin.Context) {
	id, _ := c.Get("userId")
	newId := id.(string)
	title := c.PostForm("title")
	body := c.PostForm("body")
	auth := c.PostForm("access")
	access, _ := strconv.Atoi(auth)
	helper.VerifyEmptyString(title, body, c)

	stmt, err := Db.Prepare(fmt.Sprintf("INSERT INTO posts(id, title, boby, time_created, user_id, access) VALUES(?, ?, ?,?, ?,?)"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	var post = models.Post{
		Id:          uuid.New().String(),
		Title:       title,
		Body:        body,
		Access:      access,
		TimeCreated: time.Now().Format(time.RFC850),
		UserId:      newId,
	}

	res, err := stmt.Exec(post.Id, post.Title, post.Body, post.TimeCreated, post.UserId, post.Access)
	if err != nil {
		log.Println(err.Error())
		return
	}
	total, _ := res.RowsAffected()
	if total < 1 {
		log.Println(err.Error())
		return
	}
	c.Redirect(302, "/post/home")
}

func GetPost(c *gin.Context) {
	rows, err := Db.Query("SELECT posts.id, posts.title, posts.boby, posts.time_created, posts.user_id, posts.access, users.first_name, users.last_name FROM posts INNER JOIN users ON posts.user_id = users.id;")
	if err != nil {
		log.Println(err.Error())
		c.Status(500)
		return
	}
	defer rows.Close()
	var row []models.Post
	for rows.Next() {
		var r models.Post
		err = rows.Scan(&r.Id, &r.Title, &r.Body, &r.TimeCreated, &r.UserId, &r.Access, &r.FirstName, &r.LastName)
		if err != nil {
			log.Println(err.Error())
			c.Status(500)
			return
		}
		row = append(row, r)
	}
	c.HTML(200, "Home.html", row)
}

var m = map[string]interface{}{}

func VeiwPost(c *gin.Context) {
	id := c.Param("Id")
	var p models.Post
	Db.QueryRow("SELECT * FROM posts WHERE id= ?", id).Scan(&p.Id, &p.Title, &p.Body, &p.TimeCreated, &p.UserId, &p.Access)
	m["post"] = p

	var comments []models.Comment

	rows, err := Db.Query("select comments.id, comments.commentt, comments.user_id, comments.post_id, comments.time_posted, users.first_name, users.last_name from comments inner join users on comments.user_id = users.id where comments.post_id =?", id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for rows.Next() {
		var c models.Comment
		err = rows.Scan(&c.Id, &c.Comments, &c.UserId, &c.PostId, &c.TimeCreated, &c.FirstName, &c.LastName)
		if err != nil {
			log.Println(err.Error())
			return
		}
		comments = append(comments, c)

	}
	m["comment"] = comments
	c.HTML(200, "veiwpost.html", m)
}

func UserPage(c *gin.Context) {
	id, _ := c.Get("userId")
	newId := id.(string)

	result, err := Db.Query("SELECT * FROM posts WHERE user_id = ?", newId)
	defer result.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	var results []models.Post
	for result.Next() {
		var p models.Post
		err = result.Scan(&p.Id, &p.Title, &p.Body, &p.TimeCreated, &p.UserId, &p.Access)
		if err != nil {
			log.Println(err.Error())
			return
		}
		results = append(results, p)
	}
	c.HTML(200, "userpage.html", results)
}

func DeletePost(c *gin.Context) {
	id := c.Param("Id")

	deletePost, err := Db.Prepare("DELETE FROM posts WHERE id= ?")
	defer deletePost.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	deletePost.Exec(id)
	c.Redirect(302, "/post/user")
}

func EditPost(c *gin.Context) {
	id := c.Param("Id")
	var e models.Post
	Db.QueryRow("SELECT * FROM posts WHERE id=?", id).Scan(&e.Id, &e.Title, &e.Body, &e.TimeCreated, &e.UserId, &e.Access)
	c.HTML(200, "edit.post.html", e)
}

func EditPostProcess(c *gin.Context) {
	id := c.Param("Id")

	title := c.PostForm("title")
	body := c.PostForm("body")
	auth := c.PostForm("access")
	acces, _ := strconv.Atoi(auth)
	helper.VerifyEmptyString(title, body, c)

	stmt, err := Db.Prepare("UPDATE posts SET title=?, boby=? , access= ? WHERE id=?")
	if err != nil {
		log.Println(err.Error())
		return
	}
	stmt.Exec(title, body, acces, id)
	c.Redirect(302, "/post/user")
}

func AddComment(c *gin.Context) {
	postId := c.Param("Id")
	auth, _ := c.Get("userId")
	Id := auth.(string)
	comment := c.PostForm("comment")

	stmt, err := Db.Prepare("INSERT INTO comments (id, commentt, user_id, post_id, time_posted) VALUE (?,?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
		return
	}
	com := &models.Comment{
		Id:          uuid.New().String(),
		Comments:    comment,
		UserId:      Id,
		PostId:      postId,
		TimeCreated: time.Now().Format(time.ANSIC),
	}

	stmt.Exec(com.Id, com.Comments, com.UserId, com.PostId, com.TimeCreated)
	c.Redirect(302, "/post/home")

}
