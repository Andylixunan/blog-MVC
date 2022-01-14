package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteArticleGet(c *gin.Context) {
	_, isLogin := GetLoginUsername(c)
	if !isLogin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	utils.Logger.Printf("deleting article, ID: %v", id)
	_, err = models.DeleteArticleWithID(id)
	if err != nil {
		utils.Logger.Fatalf("error occured when deleting article with ID: %v", id)
	}
	c.Redirect(http.StatusFound, "/")
}
