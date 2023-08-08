package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"user_video/model"
	"user_video/sql"
	"user_video/util"

	"log"
)

type VideoService struct {
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

func generateId(in uint) (id string) {
	id = strconv.Itoa(int(in))
	return fmt.Sprintf(id, time.Now().UnixNano())
}

func (v *VideoService) Clip_Video(c context.Context, userId uint, startTime, endTime string) (err error) {
	t := time.Now()
	video_id := generateId(userId)
	video := model.Video{
		userId,
		video_id,
		&t,
		nil,
	}
	sql_instance := sql.GetInstance()
	if err := sql_instance.Create(&video).Error; err != nil {
		log.Println("error Failed to save clip request")
		return err
	}
	//后端处理fmmpeg
	go func() {
		if err := util.ClipVideo(startTime, endTime); err != nil {
			log.Printf("clip video")
		}
	}()

	return nil
}
