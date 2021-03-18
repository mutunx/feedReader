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

func Run() {
	// 获取源文件中的订阅源
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatalf("get feed error %s", err.Error())
	}

	var results []*Result
	// 遍历订阅源 根据类型分配不同的matcher
	for _, feed := range feeds {
		log.Printf("get feed %s url %s", feed.Site, feed.Link)

		// 判断类型是否存在 不存在则用默认的
		var matcher Matcher
		if _, exist := matchers[feed.Type]; !exist {
			matcher = matchers["default"]
		} else {
			matcher = matchers[feed.Type]
		}

		// 获取结果
		result, err := matcher.Search(feed.Link)
		if err != nil {
			log.Fatalf("get result error %s", err.Error())
		}
		log.Printf("get %d's info", len(result))

		// 聚合结果
		results = append(results, result...)
	}

	// 调用matcher中的搜索方法

	// 展示结果
	Display(results)
}

func Display(results []*Result) {
	for _, item := range results {
		log.Println(item.Title)
		log.Println(item.Link)
	}
}
