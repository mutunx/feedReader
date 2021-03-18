package search

import (
	"fmt"
	"log"
)

// 解析器列表
var matchers = make(map[string]Matcher)

// 放置解析器
func Register(matcherName string, matcher Matcher) {
	// 存在则提示
	if _, exist := matchers[matcherName]; exist {
		fmt.Print("match exist")
	} else {
		// 存入
		matchers[matcherName] = matcher
	}

}

func Run(t string) {
	// 判断类型是否存在 不存在则用默认的
	var matcher Matcher
	if _, exist := matchers[t]; !exist {
		matcher = matchers["default"]
	} else {
		matcher = matchers[t]
	}

	// 获取结果
	results, err := matcher.Search()
	if err != nil {
		fmt.Errorf("get result error %s", err)
	}
	// 展示结果
	Display(results)
}

func Display(results []*Result) {
	for _, item := range results {
		log.Println(item.Title)
		log.Println(item.Link)
	}
}
