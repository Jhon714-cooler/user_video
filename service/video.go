package service

import (
	"context"
	"time"
	"user_video/api"
	"user_video/sql"
	"user_video/util"

	"log"
)

type ClipRequest struct {
	UserId    uint   `json:"user_id" binding:"required"`
	VideoURL  string `json:"video_url" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}

type Video struct {
	UserID     uint       `gorm:"column:UserID" json:"user_id"`
	VideoID    string     `gorm:"column:primaryKey;VideoID" json:"video_id"`
	UploadTime *time.Time `gorm:"column:UploadTime" json:"upload_time"` //上传时间
	FinishTime *time.Time `gorm:"colum:FinishTime" json:"finish_time"`  //剪辑结束时间
	StartTime  string     `gorm:"colum:StartTime" json:"start_time"`
	EndTime    string     `gorm:"colum:EndTime" json:"end_time"`
}

func time_parse(in_time string) (t time.Time) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, in_time)
	if err != nil {
		log.Println("parese time err", err)
		return t
	}
	return t
}

func (v *Video) New_clip(r *api.Req_clipVideo) *Video {
	UploadTime := time.Now()
	return &Video{
		r.UserId,
		"生成id",
		&UploadTime,
		nil,
		r.StartTime,
		r.EndTime,
	}
}

func (v *Video) Clip_Video(c context.Context) (err error) {
	if err := sql.Mysql_cli.Create(&v).Error; err != nil {
		log.Println("error Failed to save clip request")
		return err
	}
	//后端处理fmmpeg
	go func() {
		if err := util.ClipVideo(v.StartTime, v.EndTime); err != nil {
			log.Printf("clip video")
		}
	}()

	return nil
}
