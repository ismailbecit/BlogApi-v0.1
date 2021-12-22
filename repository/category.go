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
