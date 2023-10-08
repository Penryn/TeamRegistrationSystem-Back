package models

type User struct {
	UserID     int      `json:"-" gorm:"primaryKey"`
	Name       string   `json:"name"`
	Phone      string   `json:"phone"`
	Email      string   `json:"-"`
	Password   []byte   `json:"-"`
	Permission int      `json:"-"` //0代表用户，1代表管理员
	TeamID     int      `json:"team_id"`
	Code       string   `json:"-"`
	Team       Team     `json:"-"`
	Userinfo   Userinfo `json:"-"`
}
