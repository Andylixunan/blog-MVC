package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	pageString, ok := c.GetQuery("page")
	if !ok {
		c.Redirect(http.StatusFound, "/?page=1")
		return
	}
	page, err := strconv.Atoi(pageString)
	if page <= 0 || err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	articleList, err := models.FindArticleWithPage(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	homeBlocks := models.MakeHomeBlocks(articleList, isLogin)
	homeFooterPageCode := models.GetHomeFooterPageCode(page)
	c.HTML(http.StatusOK, "home.html", gin.H{"isLogin": isLogin, "homeBlocks": homeBlocks, "PageCode": homeFooterPageCode})
}

func GetLoginUsername(c *gin.Context) (string, bool) {
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	utils.Logger.Printf("login user: %v", loginUser)
	if loginUser == nil {
		return "", false
	}
	username := loginUser.(string)
	return username, true
}
