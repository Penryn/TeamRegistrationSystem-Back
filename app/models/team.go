package models

type Team struct {
	ID           int    `json:"id"  gorm:"foreignkey:TeamID"`
	TeamName     string `json:"team_name"`
	TeamPassword string `json:"-"`
	CaptainID    int    `json:"captain_id"`
	Slogan       string `json:"slogan"`
	Avatar       string `json:"avatar"`
	Users        []User  `json:"users"`
}
