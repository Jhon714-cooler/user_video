package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"user_video/model"
)

var instance *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/video_clip_api?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panicln("mysql client err :", err)
	}
	instance = db
	db.AutoMigrate(&model.Video{})
}

func GetInstance() *gorm.DB {
	return instance
}
