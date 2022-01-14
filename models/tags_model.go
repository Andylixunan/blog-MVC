package models

import "strings"

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
