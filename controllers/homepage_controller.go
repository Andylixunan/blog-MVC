package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	isLogin := GetSession(c)
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
		c.Redirect(http.StatusFound, fmt.Sprintf("/?page=%d", page))
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

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	utils.Logger.Printf("login user: %v", loginUser)
	return loginUser != nil
}
