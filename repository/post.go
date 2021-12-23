package repository

import (
	"blogapi/api/modal"

	"gorm.io/gorm"
)

func (rootRepo *Repositories) Post() PostRepo {
	return PostRepo{db: rootRepo.Db}
}

type PostRepo struct {
	db *gorm.DB
}

func (pt PostRepo) New(post modal.Post) error {
	err := pt.db.Debug().Create(&post)
	return err.Error
}

func (pt PostRepo) Del(id uint) error {
	u := modal.Post{}
	err := pt.db.Debug().Delete(&u, id)
	return err.Error
}

func (pt PostRepo) Query(post modal.Post, id uint) int64 {
	var count int64
	pt.db.Debug().Where("id = ? ", id).Find(&post).Count(&count)
	return count

}

func (pt PostRepo) List() []modal.Post {
	var post []modal.Post
	pt.db.Debug().Preload("Category").Preload("User").Find(&post)
	return post
}

func (pt PostRepo) CategoryFK(categoryfk uint) int64 {
	var count int64
	var post modal.Post
	pt.db.Debug().Where("categoryfk = ?", categoryfk).Find(&post).Count(&count)
	return count
}
