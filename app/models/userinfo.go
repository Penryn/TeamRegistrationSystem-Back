package models

type Userinfo struct {
	ID       int    `json:"-" `
	UserID   int    `json:"-"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Motto    string `json:"motto"`
	Avatar   string `json:"avatar"`
}
