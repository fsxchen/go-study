package model

import "github.com/jinzhu/gorm"

type UserRole struct {
	gorm.Model
	UserID string `gorm:"column:user_id;size:36;index;"` // 用户内码
	RoleID string `gorm:"column:role_id;size:36;index;"` // 角色内码
}
