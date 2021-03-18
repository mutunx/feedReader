package search

type DefaultMatcher struct{}

func (m DefaultMatcher) Search() ([]*Result, error) {
	return nil, nil
}

func init() {
	var matcher DefaultMatcher
	Register("default", matcher)
}
