package dao

import (
	"blog/dto"
	"blog/model"

	"github.com/jinzhu/gorm"
)

type User struct {
}

func (User) get(id string) model.User {
	var user model.User
	model.DB.Where("id = ?", id).First(&user)
	return user
}

// user list
func (User) List(listDto dto.GeneralListDto) (users []model.User, total uint64) {
	db = model.DB
	db.Find(&users)
	db.Model(&model.User{}).Count(&total)
	return
}

func (u User) Create(user *model.User) *gorm.DB {
	db := model.DB
	return db.Create(user)
}
