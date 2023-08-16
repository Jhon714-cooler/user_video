package main

import (
	"user_video/router"
	"user_video/service"
)

func main() {
	service.Init_cliper()
	router.Routes()
}
