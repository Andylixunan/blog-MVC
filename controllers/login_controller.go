package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, ",password:", password)
	id, err := models.QueryUserWithParam(username, utils.MD5(password))
	if id > 0 && err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}
