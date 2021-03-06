package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddArticleGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	if !isLogin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.HTML(http.StatusOK, "write_article.html", gin.H{"isLogin": isLogin})
}

func AddArticlePost(c *gin.Context) {
	//获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	utils.Logger.Printf("title: %s, tags: %s\n", title, tags)
	username, _ := GetLoginUsername(c)
	article := models.Article{
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     username,
		CreateTime: time.Now().Unix(),
	}
	_, err := models.AddArticle(article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "ok"})
	}
}
