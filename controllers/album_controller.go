package controllers

import (
	"blogweb_gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlbumGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)

	c.HTML(http.StatusOK, "album.html", gin.H{"isLogin": isLogin})
}

func AlbumPost(c *gin.Context) {
	utils.Logger.Println("start uploading file")
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": new(error)})
}
