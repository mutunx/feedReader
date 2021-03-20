package search

import "log"

type Result struct {
	Title       string
	Description string
	Link        string
}

type Matcher interface {
	Search(link string, searchItem string) ([]*Result, error)
}

// 获取源和源对应的方法
func Match(feed *Feed, matcher Matcher, results chan *Result, searchItem string) {
	// 获取结果
	result, err := matcher.Search(feed.Link, searchItem)
	if err != nil {
		log.Fatalf("get result error %s", err.Error())
	}
	log.Printf("get %d's info", len(result))

	// 读取一条后就展示一条
	for _, item := range result {
		results <- item
	}

}
