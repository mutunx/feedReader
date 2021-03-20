package search

import "log"

type Result struct {
	Title       string
	Description string
	Link        string
}

type Matcher interface {
	Search(feed *Feed, searchItem string) ([]*Result, error)
}

// 获取源和源对应的方法
func Match(feed *Feed, matcher Matcher, results chan *Result, searchItem string) {
	// 获取结果
	result, err := matcher.Search(feed, searchItem)
	if err != nil {
		log.Println("get result error ", err.Error())
	}

	// 读取一条后就展示一条
	for _, item := range result {
		results <- item
	}

}

func Display(results chan *Result) {
	for item := range results {
		log.Println(item.Title)
		log.Println(item.Link)
	}
}
