package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddArticleGet(c *gin.Context) {
	isLogin := GetSession(c)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"isLogin": isLogin})
}

func AddArticlePost(c *gin.Context) {
	//获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	utils.Logger.Printf("title:%s, tags:%s\n", title, tags)
	article := models.Article{
		Title:      title,
		Short:      short,
		Content:    content,
		Author:     "Andy",
		CreateTime: time.Now().Unix(),
	}
	_, err := models.AddArticle(article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "ok"})
	}
}
