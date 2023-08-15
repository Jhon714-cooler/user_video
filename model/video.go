package model

import "time"

type Video struct {
	UserID     uint       `gorm:"column:UserID" json:"user_id"`
	VideoID    string     `gorm:"column:VideoID;primaryKey" json:"video_id"`
	UploadTime *time.Time `gorm:"column:UploadTime" json:"upload_time"` //上传时间
	FinishTime *time.Time `gorm:"colum:FinishTime" json:"finish_time"`  //剪辑结束时间
	//StartTime  string     `gorm:"colum:StartTime" json:"start_time"`
	//EndTime    string     `gorm:"colum:EndTime" json:"end_time"`
	Url string `gorm:"colum:url" json:"url"`
}
