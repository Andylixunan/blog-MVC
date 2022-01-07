package models

import (
	"blogweb_gin/database"
	"blogweb_gin/utils"
	"fmt"
)

type Article struct {
	ID         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64
}

func AddArticle(article Article) (int64, error) {
	return database.ModifyDB("insert into article(title,tags,short,content,author,createTime) values(?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime)
}

//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	utils.Logger.Println("---------->page", page)
	//从配置中获取每页的文章数量
	return QueryArticleWithPage(page, utils.ArticleDisplayNum)
}

/*
分页查询数据库
limit分页查询语句，
    语法：limit m，n

    m代表从多少位开始获取，与id值无关
    n代表获取多少条数据
*/

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createTime from article " + sql
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var articleList []Article
	for rows.Next() {
		var article Article
		rows.Scan(&article.ID, &article.Title, &article.Tags, &article.Short, &article.Content, &article.Author, &article.CreateTime)
		articleList = append(articleList, article)
	}
	return articleList, nil
}
