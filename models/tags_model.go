package models

import (
	"strings"
)

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		//普及&保险&基础知识
		// 普及、保险、基础知识
		// 客户端，get，post等方法
		//DOM&前端&xiangzai
		// DOM 、 前端 、 xiangzai
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
		// key  value
		// 普及  1
		// 保险  1
		// 基础知识 1
		// 客户端，get，post等方法  1
		// DOM   1
		// 前端   1
		// xiangzai 1
	}
	return tagsMap
}
