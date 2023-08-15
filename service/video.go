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

var (
	videoProcess = make(map[string]int)
)

type VideoService struct {
	Signal    chan struct{}
	UserId    uint
	VideoURL  string
	StartTime string
	EndTime   string
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
	random := fmt.Sprintf("%04d", time.Now().UnixNano()%1000)
	return fmt.Sprintf(id, time.Now().UnixNano(), random)
}

func (v *VideoService) Clip_Video(c context.Context, userId uint, startTime, endTime string) (video_id string, err error) {
	t := time.Now()
	video_id = generateId(userId)
	video := model.Video{
		userId,
		video_id,
		&t,
		nil,
		"",
	}
	sql_instance := sql.GetInstance()
	if err := sql_instance.Create(&video).Error; err != nil {
		log.Println("error Failed to save clip request")
		return "", err
	}
	//后端处理fmmpeg
	go func() {
		videoProcess[video_id] = -1
		if err := util.ClipVideo(startTime, endTime); err != nil {
			videoProcess[video_id] = 0
			log.Printf("clip video fail", err)
		}
		videoProcess[video_id] = 1
	}()
	defer log.Printf("FINISH VIDEO")
	return video_id, nil
}
func (v *VideoService) GetVideoResult(video_id string) (url string, err error) {
	var video model.Video
	sql_instance := sql.GetInstance()
	sql_instance.Where("VideoID = ?", video_id).Find(video)

	return video.Url, err
}
