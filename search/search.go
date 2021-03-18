package search

import "fmt"

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

	matcher.Search()

}
