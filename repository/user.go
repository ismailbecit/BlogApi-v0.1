package repository

import (
	"blogapi/api/modal"

	"gorm.io/gorm"
)

func (rootRepo *Repositories) User() UserRepo {
	return UserRepo{db: rootRepo.Db}
}

type UserRepo struct {
	db *gorm.DB
}

func (us UserRepo) New(user modal.User) modal.User {
	us.db.Debug().Create(&user)
	return user
}

func (us UserRepo) EmailQuery(email string) int64 {
	var count int64
	var user modal.User
	us.db.Debug().Where("email = ?", email).Find(&user).Count(&count)
	return count

}
func (us UserRepo) UserInfo(id uint) modal.User {
	user := modal.User{}
	us.db.Debug().Find(&user, id)
	return user
}

func (us UserRepo) UserList(user []modal.User) []modal.User {
	us.db.Debug().Find(&user)
	return user
}
func (us UserRepo) UserDel(id uint) error {
	err := us.db.Debug().Delete(&modal.User{}, id)
	return err.Error
}
