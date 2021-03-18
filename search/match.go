package search

type Result struct {
	Title       string
	Description string
	Link        string
}

type Matcher interface {
	Search(link string) ([]*Result, error)
}
