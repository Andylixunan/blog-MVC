package models

import (
	"blogweb_gin/database"
	"blogweb_gin/utils"
	"strings"
)

func HandleTagsList(tags []string) map[string]int {
	tagsMap := make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, v := range tagList {
			tagsMap[v]++
		}
	}
	return tagsMap
}

func QueryAllTags() []string {
	rows, err := database.QueryDB("select tags from article")
	if err != nil {
		utils.Logger.Fatalf("querying tags failed --> err: %v", err)
	}
	tagList := []string{}
	for rows.Next() {
		tag := ""
		rows.Scan(&tag)
		tagList = append(tagList, tag)
	}
	return tagList
}

func FindArticleWithTag(tag string) ([]Article, error) {
	condition := " where tags like '%&" + tag + "&%'"
	condition += " or tags like '%&" + tag + "'"
	condition += " or tags like '" + tag + "&%'"
	condition += " or tags like '" + tag + "'"
	return QueryArticleWithCon(condition)
}
