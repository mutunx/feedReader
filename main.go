package main

import (
	_ "feedReader/matchers"
	"feedReader/search"
	"log"
	"os"
)

func init() {
	// 设置日志展示位置
	log.SetOutput(os.Stdout)
}

func main() {
	// 方法主入口
	search.Run("为什么")
}
