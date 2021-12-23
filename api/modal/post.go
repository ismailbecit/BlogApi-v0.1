package modal

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string   `json:"title"`
	Content    string   `json:"context"`
	User       User     `gorm:"foreignkey:userfk"`
	Userfk     uint     `gorm:"column:userfk;" json:"userfk"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignkey:categoryfk"`
	Categoryfk uint     `gorm:"column:categoryfk" json:"categoryfk"`
}
