package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	// gorm.Model
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	// ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := DB.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	// if tag.ID > 0 {
	return true, nil
	// }

	// return false, nil
}

// AddTag Add a Tag
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := DB.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	// if pageSize > 0 && pageNum > 0 {
	// 	fmt.Println("nnnn")
	// 	err = DB.Model(&Tag{}).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	// } else {
	// 	fmt.Println("not")
	// }

	err = DB.Find(&tags).Error

	fmt.Println(err)
	fmt.Println(len(tags))

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal counts the total number of tags based on the constraint
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := DB.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		fmt.Println(err)
		return 0, err
	}
	return count, nil
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := DB.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	// if tag.ID > 0 {
	return true, nil
	// }

	// return false, nil
}

// DeleteTag delete a tag
func DeleteTag(id int) error {
	if err := DB.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// EditTag modify a single tag
func EditTag(id int, data interface{}) error {
	if err := DB.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllTag clear all tag
func CleanAllTag() (bool, error) {
	if err := DB.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
