package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type (
	// todoModel 包括了 todoModel 的字段类型
	User struct {
		gorm.Model
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// transformedTodo 代表格式化的 todo 结构体
)

func (user *User) Save() error {
	return DB.Create(user).Error
}

func (user *User) ListAccount() (users []*User, err error) {
	err = DB.Find(&users).Error
	return
}

type UserRole struct {
	gorm.Model
	UserID string `gorm:"column:user_id;size:36;index;"` // 用户内码
	RoleID string `gorm:"column:role_id;size:36;index;"` // 角色内码
}
