package models

import "blogweb_gin/database"

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
