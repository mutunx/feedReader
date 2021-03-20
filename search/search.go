package search

import (
	"log"
	"sync"
)

// 解析器列表 用于解析器的获取和注册
var matchers = make(map[string]Matcher)

// 解析器注册
func Register(matcherName string, matcher Matcher) {
	// 解析器如果已注册则报错
	if _, exist := matchers[matcherName]; exist {
		log.Fatalf("mathers is already register")
	}
	// 注册
	log.Println("register ", matcherName, "Mather")
	matchers[matcherName] = matcher

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
		matcher, exist := matchers[feed.Type]
		if !exist {
			matcher = matchers["default"]
		}

		go func(feed *Feed, matcher Matcher) {
			Match(feed, matcher, results, searchItem)
			waitGroup.Done()
		}(feed, matcher)
	}
	// go range 遍历channel时 如果channel关闭就会自动退出
	go func() {
		waitGroup.Wait()
		close(results)
	}()

	// 展示结果
	Display(results)

}
