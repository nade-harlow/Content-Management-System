package models

type Comment struct {
	Id          string `json:"id"`
	Comments    string `json:"comments"`
	UserId      string `json:"user_id"`
	PostId      string `json:"post_id"`
	TimeCreated string `json:"time_created"`
}
