package models

type Team struct {
	ID           int    `json:"id"  gorm:"foreignkey:TeamID"`
	TeamName     string `json:"team_name"`
	TeamPassword string `json:"-"`
	CaptainID    int    `json:"-"`
	CaptainName  string `json:"captain_name"`
	Slogan       string `json:"slogan"`
	Avatar       string `json:"avatar"`
	Confirm      int    `json:"confirm"`
	Number       int    `json:"number"`
	Users        []User `json:"users"`
}
