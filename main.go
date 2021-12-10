package main

import (
	"blogweb_gin/database"
	"blogweb_gin/routers"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()
	//静态资源
	router.Static("/static", "./static")
	err := router.Run(":80")
	if err != nil {
		return
	}
}
