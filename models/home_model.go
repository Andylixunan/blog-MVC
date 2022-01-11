package models

import (
	"blogweb_gin/database"
	"blogweb_gin/utils"
	"fmt"
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

// HomeFooterPageCode 用来记录分页
type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

var totalArticleNumbers int

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
			Link:       "/article/show/" + strconv.Itoa(article.ID),
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
		tagLink = append(tagLink, TagLink{tag, "/article/tags/?tag=" + tag})
	}
	return tagLink
}

//-----------翻页-----------
//page是当前的页数
func GetHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	totalArticleNum, err := GetTotalArticleNums()
	if err != nil {
		utils.Logger.Fatal(err)
	}
	totalPageNums := (totalArticleNum-1)/utils.ArticleDisplayNum + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, totalPageNums)
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	if page >= totalPageNums {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	if pageCode.HasPre {
		pageCode.PreLink = fmt.Sprintf("/?page=%d", page-1)
	}
	if pageCode.HasNext {
		pageCode.NextLink = fmt.Sprintf("/?page=%d", page+1)
	}
	return pageCode
}

func GetTotalArticleNums() (int, error) {
	if totalArticleNumbers == 0 {
		num, err := QueryTotalArticleNums()
		if err != nil {
			return 0, err
		}
		totalArticleNumbers = num
	}
	return totalArticleNumbers, nil
}

func ResetTotalArticleNums() error {
	num, err := QueryTotalArticleNums()
	if err != nil {
		return err
	}
	totalArticleNumbers = num
	return nil
}

func QueryTotalArticleNums() (int, error) {
	num := 0
	row := database.QueryRowDB("select count(id) from article")
	err := row.Scan(&num)
	if err != nil {
		return 0, err
	}
	return num, nil
}
