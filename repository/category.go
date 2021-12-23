package repository

import (
	"blogapi/api/modal"

	"gorm.io/gorm"
)

func (rootRepo *Repositories) Category() CategoryRepo {
	return CategoryRepo{db: rootRepo.Db}
}

type CategoryRepo struct {
	db *gorm.DB
}

// veri tabanı katmanında sadice vt işlemleri yap controller içinde buraya vt nesnesi gönder başka bir işlem yaptırma !!!
func (ct CategoryRepo) New(category modal.Category) error {
	err := ct.db.Create(&category)
	return err.Error
}

func (ct CategoryRepo) Del(category modal.Category, id uint) error {

	err := ct.db.Unscoped().Delete(&category, id)
	return err.Error
}

func (ct CategoryRepo) Query(category modal.Category, id uint) int64 {
	var count int64
	ct.db.Where("id = ? ", id).Find(&category).Count(&count)
	return count

}
