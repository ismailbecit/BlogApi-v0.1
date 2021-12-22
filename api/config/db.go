package config

import (
	"blogapi/api/modal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	Database = Conn()
	Migration()
}

func Conn() *gorm.DB {
	dsn := "root:aea76026@tcp(127.0.0.1:3306)/blogapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Veri Tabanına Bağlanılamadı")
	}
	return db
}

func Migration() {
	err := Database.AutoMigrate(
		&modal.User{},
		&modal.Category{},
		&modal.Post{},
	)
	if err != nil {
		panic("Tablolar Düzgün OLuşturulamadı!")
	}
}
