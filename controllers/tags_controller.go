package controllers

import (
	"blogweb_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TagsGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	tags := models.QueryAllTags()
	tagsMap := models.HandleTagsList(tags)
	c.HTML(http.StatusOK, "tags.html", gin.H{"Tags": tagsMap, "isLogin": isLogin})
}
