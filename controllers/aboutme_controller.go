package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutMeGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	c.HTML(http.StatusOK, "aboutme.html", gin.H{
		"isLogin": isLogin,
		"wechat":  "Wechat: lxnzqforever",
		"tel":     "Tel: (+852) 52628117",
		"name":    "LI Xunan",
	})
}
