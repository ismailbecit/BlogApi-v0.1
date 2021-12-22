package repository

import (
	"blogapi/api/modal"
	"blogapi/request"

	"gorm.io/gorm"
)

func (rootRepo *Repositories) Category() CategoryRepo {
	return CategoryRepo{db: rootRepo.Db}
}

type CategoryRepo struct {
	db *gorm.DB
}

func (ct CategoryRepo) New(rq *request.CategoryInsert) error {
	err := ct.db.Create(modal.Category{Name: rq.Name})
	return err.Error
}
