package router

import (
	"github.com/gin-gonic/gin"
	"user_video/api"
)

func Routes() {
	port := ":9091"

	r := gin.Default()

	r.GET("/video/clip/result", api.GetVideo)
	r.POST("/video/upload", api.UploadVideo)

	r.Run(port)
}
