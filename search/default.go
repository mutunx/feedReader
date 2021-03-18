package search

import "fmt"

type DefaultMatcher struct{}

func (m DefaultMatcher) Search() {
	fmt.Print("this is default matcher\n")
}

func init() {
	var matcher DefaultMatcher
	Register("default", matcher)
}
