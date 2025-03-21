package main

import (
	"os"

	_ "go.uber.org/automaxprocs" // 使程序自动设置 GOMAXPROCS 以匹配 Linux 容器 CPU 配额

	"github.com/herman7/hermanblogserver/internal/hermanblog"
)

// Default entry (main function)
func main() {
	command := hermanblog.NewHermanBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(2)
	}
}
