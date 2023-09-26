package models

type User struct {
	UserID   int   `json:"user_id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Email    string `json:"email"` 
	Password []byte `json:"-"`
	Type     int8   `json:"type"`
}
