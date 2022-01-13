package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateArticleGet(c *gin.Context) {
	isLogin := CheckLogin(c)
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	utils.Logger.Printf("showing update page for article: %v", id)
	articles, err := models.QueryArticleWithCon(fmt.Sprintf("where id = %d", id))
	if err != nil {
		utils.Logger.Fatalf("query articles failed --> error: %v", err)
	}
	article := articles[0]
	c.HTML(http.StatusOK, "write_article.html", gin.H{
		"isLogin": isLogin,
		"Title":   article.Title,
		"Tags":    article.Tags,
		"Short":   article.Short,
		"Content": article.Content,
		"ID":      article.ID,
	})
}

func UpdateArticlePost(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	utils.Logger.Printf("doing update for article: %v", id)
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	_, err = models.UpdateArticle(models.Article{
		ID:      id,
		Title:   title,
		Tags:    tags,
		Short:   short,
		Content: content,
	})
	response := gin.H{"code": 1, "message": "更新成功"}
	if err != nil {
		response = gin.H{"code": 0, "message": "更新失败"}
	}
	c.JSON(http.StatusOK, response)
}
