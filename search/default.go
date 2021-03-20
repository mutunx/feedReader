package search

type DefaultMatcher struct{}

func (m DefaultMatcher) Search(feed *Feed, searchItem string) ([]*Result, error) {
	return nil, nil
}

func init() {
	var matcher DefaultMatcher
	Register("default", matcher)
}
