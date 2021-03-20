package search

import (
	"fmt"
	"log"
	"sync"
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

func Run(searchItem string) {
	// 获取源文件中的订阅源
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatalf("get feed error %s", err.Error())
	}

	// 监听器
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	results := make(chan *Result)
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

		go func(feed *Feed, matcher Matcher, results chan *Result) {
			Match(feed, matcher, results, searchItem)
			waitGroup.Done()
		}(feed, matcher, results)
	}
	// go range 遍历channel时 如果channel关闭就会自动退出
	go func() {
		waitGroup.Wait()
		close(results)
	}()

	// 展示结果
	Display(results)

}

func Display(results chan *Result) {
	for item := range results {
		log.Println(item.Title)
		log.Println(item.Link)
	}
}
