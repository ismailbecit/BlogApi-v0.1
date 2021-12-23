package modal

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Admin    bool   `gorm:"default:false" json:"admin"`
}
