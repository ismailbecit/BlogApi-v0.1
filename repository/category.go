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
	err := ct.db.Debug().Create(&category)
	return err.Error
}

func (ct CategoryRepo) List() []modal.Category {
	var category []modal.Category
	ct.db.Debug().Find(&category)
	return category
}

func (ct CategoryRepo) Del(id uint) error {
	u := modal.Category{}
	err := ct.db.Debug().Where("id = ?", id).Delete(&u)
	return err.Error
}

func (ct CategoryRepo) Query(id uint) int64 {
	var count int64
	u := modal.Category{}
	ct.db.Debug().Where("id = ?", id).Find(&u).Count(&count)
	return count

}
