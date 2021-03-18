package matchers

import (
	"feedReader/search"
	"fmt"
)

type rssMatcher struct{}

func (r rssMatcher) Search() {
	fmt.Print("this is rss matcher\n")
}

func init() {
	var r rssMatcher
	search.Register("rss", r)
}
