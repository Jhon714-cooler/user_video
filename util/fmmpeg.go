package util

import (
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"strconv"
	"sync"
)

var mutex sync.Mutex

// clipVideo clips a video according to the clip request.
func ClipVideo(start, end string) error {
	mutex.Lock()
	defer mutex.Unlock()
	var infile, opfile string = "E:\\go\\bin\\src\\user_video\\util\\tests.mp4", "E:\\go\\bin\\src\\user_video\\util\\out2.mp4"
	//var abdpath = "E:\\go\\bin\\src\\user_video\\util"
	ss, _ := strconv.Atoi(start)
	tt, _ := strconv.Atoi(end)

	err := ffmpeg_go.Input(infile, ffmpeg_go.KwArgs{"ss": ss}).
		Output(opfile, ffmpeg_go.KwArgs{"t": tt}).OverWriteOutput().Run()
	if err != nil {
		log.Println("剪辑失败", err)
		return err
	}
	return nil
}
