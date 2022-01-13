package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_, err := models.QueryUserWithParam(username, utils.MD5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	} else {
		session := sessions.Default(c)
		session.Set("login_user", username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登录成功"})
	}
}
