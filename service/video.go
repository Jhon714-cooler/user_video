package service

import (
	"context"
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"runtime"
	"strconv"
	"sync"
	"time"
	"user_video/model"
	"user_video/sql"

	"log"
)

var (
	VideoProcess = sync.Map{}
)

type VideoService struct {
	Signal    chan struct{}
	UserId    uint
	VideoURL  string
	StartTime string
	EndTime   string
}

func generateId(in uint) (id string) {
	in_id := strconv.Itoa(int(in))
	timestamp := time.Now().Unix() % 100000000 // 获取当前时间戳（以秒为单位）
	id = fmt.Sprintf("%s%08d", in_id, timestamp)
	return id
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
	Cliper_server.C <- startTime
	Cliper_server.C <- endTime
	Cliper_server.C <- video_id
	VideoProcess.Store(video_id, 0)
	//go func() {
	//	videoProcess[video_id] = -1
	//	if err := util.ClipVideo(startTime, endTime); err != nil {
	//		videoProcess[video_id] = 0
	//		log.Printf("clip video fail %v", err)
	//	}
	//	videoProcess[video_id] = 1
	//}()
	//defer log.Printf("FINISH VIDEO")
	return video_id, nil
}
func (v *VideoService) GetVideoResult(video_id string) (url string, err error) {
	var video model.Video
	sql_instance := sql.GetInstance()
	sql_instance.Where("VideoID = ?", video_id).Find(video)

	return video.Url, err
}

var Cliper_server Cliper

type Cliper struct {
	C       chan string
	start   string
	end     string
	videoId string
}

func Init_cliper() {
	go Cliper_server.Newcliper()
}
func (c *Cliper) Newcliper() {
	c.C = make(chan string, 3)
	for {
		var ok bool
		c.start, ok = <-c.C
		c.end, ok = <-c.C
		c.videoId = <-c.C
		if !ok {
			log.Println("cliper channel err ")
		}
		if err := c.ClipVideo(); err != nil {
			VideoProcess.Store(c.videoId, -1)
			_, file, line, ok := runtime.Caller(0)
			if ok {
				log.Printf("Error at %s:%d - %v", file, line, err)
			} else {
				log.Printf("Error: %v", err)
			}
		}
		VideoProcess.Store(c.videoId, 1)
		//_, _ := VideoProcess.Load(c.videoId)
		t := time.Now()
		video_tabel := sql.GetInstance().Model(&model.Video{})

		result := video_tabel.Where("videoId = ?", c.videoId).Updates(map[string]interface{}{
			"finish_time": t,
			"url":         "www.example.com 返回url",
		})
		if result.Error != nil {
			log.Println(result.Error)
		}
	}
}

// clipVideo clips a video according to the clip request.
func (c *Cliper) ClipVideo() error {

	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	var infile, opfile string = "E:\\go\\bin\\src\\user_video\\util\\tests.mp4", "E:\\go\\bin\\src\\user_video\\util\\out2.mp4"
	//var abdpath = "E:\\go\\bin\\src\\user_video\\util"
	ss, _ := strconv.Atoi(c.start)
	tt, _ := strconv.Atoi(c.end)

	err := ffmpeg_go.Input(infile, ffmpeg_go.KwArgs{"ss": ss}).
		Output(opfile, ffmpeg_go.KwArgs{"t": tt}).OverWriteOutput().Run()
	if err != nil {
		log.Println("剪辑失败", err)
		return err
	}
	return nil
}
