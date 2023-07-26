package sql

import (
	"github.com/jinzhu/gorm"
	"log"
	"user_video/service"
)

var Mysql_cli *gorm.DB

func NewClient() {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/video_clip_api?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panicln("mysql err")
	}
	Mysql_cli = db
	defer Mysql_cli.Close()
	db.AutoMigrate(&service.ClipRequest{})
}
