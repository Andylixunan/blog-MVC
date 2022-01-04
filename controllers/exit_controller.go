package controllers

import (
	"blogweb_gin/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ExitGet(c *gin.Context) {
	session := sessions.Default(c)
	utils.Logger.Printf("delete session: %v", session.Get("login_user"))
	session.Delete("login_user")
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
