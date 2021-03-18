package main

import (
	_ "feedReader/matchers"
	"feedReader/search"
)

func main() {
	search.Run("rss")
	search.Run("sdfsdf")
}
