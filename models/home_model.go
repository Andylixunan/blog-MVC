package models

import (
	"blogweb_gin/utils"
	"bytes"
	"html/template"
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
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
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

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的变量替换为传进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	utils.Logger.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
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
