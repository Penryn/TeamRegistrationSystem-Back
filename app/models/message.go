package models

type Message struct {
	ID          int    `json:"-"`
	UserID      int    `json:"user_id"`
	Information string `json:"information"`
	Time        string `json:"time"`
}
