package main

import (
	_ "feedReader/matchers"
	"feedReader/search"
)

func main() {
	search.Run("为什么")
}
