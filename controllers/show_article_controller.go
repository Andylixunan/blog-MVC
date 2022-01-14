package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// showing details of an article
func ShowArticleGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	utils.Logger.Printf("show article: %d", id)
	articles, err := models.QueryArticleWithCon(fmt.Sprintf("where id = %d", id))
	if err != nil {
		utils.Logger.Fatalf("query articles failed --> error: %v", err)
	}
	article := articles[0]
	c.HTML(http.StatusOK, "show_article.html", gin.H{"isLogin": isLogin, "Title": article.Title, "Content": utils.SwitchMarkdownToHTML(article.Content)})
}
