package models

type Message struct {
	ID          int    `json:"-"`
	UserID      int    `json:"-"`
	Information string `json:"information"`
	Time        string `json:"time"`
}
