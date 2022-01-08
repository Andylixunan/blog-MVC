package models

import (
	"blogweb_gin/utils"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string

	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

//----------首页显示内容---------

func MakeHomeBlocks(articles []Article, isLogin bool) []HomeBlockParam {
	homeBlocks := []HomeBlockParam{}
	for _, article := range articles {
		homeParam := HomeBlockParam{
			Id:         article.ID,
			Title:      article.Title,
			Short:      article.Short,
			Content:    article.Content,
			Author:     article.Author,
			Link:       "/show/" + strconv.Itoa(article.ID),
			UpdateLink: "/article/update?id=" + strconv.Itoa(article.ID),
			DeleteLink: "/article/delete?id=" + strconv.Itoa(article.ID),
			IsLogin:    isLogin,
			CreateTime: utils.SwitchTimeStampToData(article.CreateTime),
			Tags:       createTagsLinks(article.Tags),
		}
		homeBlocks = append(homeBlocks, homeParam)
	}
	return homeBlocks
}

//将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tags, "&")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}
