package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	log.Println(username, password, repassword)
	id, err := models.QueryWithUsername(username)
	if id != 0 && err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已存在"})
		return
	}
	password = utils.MD5(password)
	user := models.User{
		Username:   username,
		Password:   password,
		Status:     0,
		CreateTime: time.Now().Unix(),
	}
	_, err = models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
	}
}
