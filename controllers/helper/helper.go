package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func IsEmailValid(e string, c *gin.Context) string {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(e) == false {
		c.String(400, "Invalid email\n")
		c.Redirect(http.StatusNotModified, "/form")
		return ""
	}
	return e
}

func VerifyEmpty(a, b string, c *gin.Context) {
	if a == "" {
		c.String(400, "Invalid first name")
		return
	}
	if b == "" {
		c.String(400, "Invalid last name")
		return
	}
}
