package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func AlbumGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	images, err := models.FindAllImagesFromAlbum()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.HTML(http.StatusOK, "album.html", gin.H{"isLogin": isLogin, "Image": images})
}

func AlbumPost(c *gin.Context) {
	fileHeader, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	}
	permittedFileType := map[string]struct{}{
		".jpg":  {},
		".jpeg": {},
		".png":  {},
		".JPG":  {},
		".JPEG": {},
		".PNG":  {},
	}
	fileExtension := filepath.Ext(fileHeader.Filename)
	if _, ok := permittedFileType[fileExtension]; !ok {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "不支持的文件类型"})
		return
	}
	utils.Logger.Printf("start uploading file: %v, size: %v", fileHeader.Filename, fileHeader.Size)
	now := time.Now()
	fileDir := fmt.Sprintf("static/upload/%d/%d/%d", now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "创建文件夹失败"})
		return
	}
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePath := filepath.Join(fileDir, fileName)
	c.SaveUploadedFile(fileHeader, filePath)
	image := models.Image{
		FilePath:   filePath,
		FileName:   fileName,
		Status:     0,
		CreateTime: timeStamp,
	}
	_, err = models.InsertIntoAlbum(image)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传成功"})
}
