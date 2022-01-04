package main

import (
	"blogweb_gin/database"
	"blogweb_gin/routers"
	"blogweb_gin/utils"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()
	err := router.Run(":80")
	if err != nil {
		utils.Logger.Fatalf("failed to start server\n")
		return
	}
}
