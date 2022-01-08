package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	isLogin := GetSession(c)
	page := 1
	articleList, err := models.FindArticleWithPage(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	homeBlocks := models.MakeHomeBlocks(articleList, isLogin)
	c.HTML(http.StatusOK, "home.html", gin.H{"isLogin": isLogin, "homeBlocks": homeBlocks})
}

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	utils.Logger.Printf("login user: %v", loginUser)
	return loginUser != nil
}
