package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeGet(c *gin.Context) {
	isLogin := GetSession(c)
	c.HTML(http.StatusOK, "home.html", gin.H{"isLogin": isLogin})
}

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	log.Printf("login user: %v", loginUser)
	return loginUser != nil
}
