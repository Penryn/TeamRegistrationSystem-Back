package models

type User struct {
	UserID     int    `json:"user_id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   []byte `json:"-"`
	Permission int    `json:"-"`   //0代表用户，1代表管理员
	TeamID    int     `json:"team_id"`
	Team      Team    `json:"-"`
	Userinfo  Userinfo `json:"-"`
}
