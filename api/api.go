package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Req_clipVideo struct {
	VideoURL  string `json:"video_url" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
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
	log.Println("hello")
}
