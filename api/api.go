package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req_clipVideo struct {
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

func UploadVideo(c *gin.Context) {
	var req Req_clipVideo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "cliping video",
		"cliResult_url": "查询剪辑进度",
		"video_id":      "视频id 查询剪辑进度",
	})
}
func GetVideo(c *gin.Context) {
	var req Req_videoUrl
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          "cliped_video_url",
		"result_video_url": "最终视频url",
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
