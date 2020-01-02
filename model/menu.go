package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Menu struct {
	Model
	MenuLevel  uint   `json:"-"`
	ParentId   string `json:"parentId"`
	ParentPath string `json:"parentPath"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Hidden     bool   `json:"hidden"`
	Component  string `json:"component"`
	Router     string `json:"router"`
	Icon       string `json:"icon"`
	Meta       string `json:"meta"`
	NickName   string `json:"nickName"`
	Children   []Menu `json:"children"`
}

func (menu *Menu) Save() error {
	if err := DB.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

// GetTags gets a list of tags based on paging and constraints
func GetMenus() (menus []Menu, err error) {
	// if pageSize > 0 && pageNum > 0 {
	// 	fmt.Println("nnnn")
	// 	err = DB.Model(&Tag{}).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	// } else {
	// 	fmt.Println("not")
	// }

	err = DB.Find(&menus).Error

	fmt.Println(err)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return menus, err
}
