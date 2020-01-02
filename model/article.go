package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title           string `json:"title"`
	Desc            string `json:"desc"`
	Content         string `json:"content"`
	ContentMarkdown string `json:"content_markdown"`
	CreatedBy       string `json:"created_by"`
	ModifiedBy      string `json:"modified_by"`
	State           int    `json:"state"`
}

func ExistArticleByID(id string) bool {
	var article Article
	DB.Select("id").Where("id = ?", id).First(&article)

	// if article.ID > 0 {
	return true
	// }

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	DB.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	DB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id string) (article Article) {
	DB.Where("id = ?", id).First(&article)
	DB.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id string, data interface{}) bool {
	DB.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func (article *Article) Save() bool {
	DB.Create(article)
	return true
}

func DeleteArticle(id string) bool {
	DB.Where("id = ?", id).Delete(Article{})

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
