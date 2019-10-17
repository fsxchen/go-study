package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	RecordID string  `gorm:"column:record_id;size:36;index;"` // 记录内码
	Name     *string `gorm:"column:name;size:100;index;"`     // 角色名称
	Sequence *int    `gorm:"column:sequence;index;"`          // 排序值
	Memo     *string `gorm:"column:memo;size:200;"`           // 备注
	Creator  *string `gorm:"column:creator;size:36;"`         // 创建者
}

func (role *Role) Save() error {
	return DB.Create(role).Error
}
