package modal

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	User       User     `gorm:"foreignkey:userfk"`
	Userfk     uint     `gorm:"column:userfk" json:"userfk"`
	Category   Category `gorm:"foreignkey:categoryfk"`
	Categoryfk uint     `gorm:"column:categoryfk" json:"categoryfk"`
	Title      string   `json:"title"`
	Content    string   `json:"context"`
}
