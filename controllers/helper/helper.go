package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"regexp"
)

var DbSessions = map[string]string{}
var UserDb = map[string]string{}

func SetCookie(c *gin.Context) {
	sId := uuid.New().String()
	cookie, err := c.Cookie("")
	if err != nil {
		cookie = "not set"
		c.SetCookie("session", sId, 3600, "/user", "localhost", true, true)

	}
	log.Println("cookie value: ", cookie)

	//cc,err := r.Cookie("session")
	//if err != nil{
	//	sId := uuid.New().String()
	//	cc = &http.Cookie{
	//		Name: "session",
	//		Value: sId,
	//	}
	//	http.SetCookie(w,cc)
	//}
	//var u models.User
	//if un, ok := DbSessions[cc.Value];ok{
	//	u = UserDb[un]
	//}

}

func IsEmailValid(e string, c *gin.Context) string {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(e) == false {
		c.String(400, "Invalid email\n")
		c.Redirect(http.StatusNotModified, "/form")
		return ""
	}
	return e
}

func VerifyEmpty(fristName, lastName string, c *gin.Context) {
	if fristName == "" {
		c.String(400, "Invalid first name\n")
		return
	}
	if lastName == "" {
		c.String(400, "Invalid last name\n")
		return
	}
}

func VerifyEmptyString(title, body string, c *gin.Context) {
	if title == "" {
		c.String(400, "Enter title field\n")
		return
	}
	if body == "" {
		c.String(400, "Enter body field\n")
		return
	}
}
