package models

type User struct {
	UserID     int    `json:"user_id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Password   []byte `json:"-"`
	Permission int    `json:"-"`
	Userinfo  Userinfo `json:"user_info"`
}
