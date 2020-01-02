package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		uuid, _ := uuid.NewV4()

		scope.SetColumn("ID", uuid)
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("sqlite3", "./test.db")
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err == nil {
		DB = db
		//db.LogMode(true)
		db.AutoMigrate(&User{}, &Role{}, &UserRole{}, &Tag{}, &Article{}, &Menu{})
		// db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
		// var u User
		// u.Username = "aaa"
		// u.Password = "bbb"
		// u.Save()
		db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		return db, err
	}
	return nil, err
}
