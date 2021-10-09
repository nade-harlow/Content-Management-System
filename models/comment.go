package models

import (
	"database/sql"
)

type Comment struct {
	Id          string `json:"id"`
	Comments    string `json:"comments"`
	UserId      string `json:"user_id"`
	PostId      string `json:"post_id"`
	TimeCreated string `json:"time_created"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

var Db *sql.DB
