package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_video/service"
)

type Req_clipVideo struct {
	UserId    uint   `json:"user_id" binding:"required"`
	VideoURL  string `json:"video_url" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}
type Req_VideoStatus struct {
	Videoid string `json:"video_id" binding:"required"`
}
type Req_videoUrl struct {
	Videoid string `json:"videoid" binding:"required"`
}

var video service.VideoService

func UploadVideo(c *gin.Context) {
	var req Req_clipVideo
	var id string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	} else {
		if id, err = video.Clip_Video(c.Request.Context(), req.UserId, req.StartTime, req.EndTime); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "cliping video",
		"cliResult_url": "查询剪辑进度",
		"video_id":      id,
	})
}
func GetVideo(c *gin.Context) {
	var req Req_videoUrl
	var url string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	} else {
		url, err = video.GetVideoResult(req.Videoid)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          "cliped_video_url",
		"result_video_url": url,
	})
}
func GetVideostatus(c *gin.Context) {
	var req Req_VideoStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":           "clip_video_status",
		"clip_video_status": "获取剪辑视频进度",
	})
}
