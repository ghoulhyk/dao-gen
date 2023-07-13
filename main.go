package main

import (
	"daogen/generator"
	"flag"
)

var (
	basicPath = flag.String("d", "", "dao根目录")
	confFile  = flag.String("c", "", "配置文件")
)

func main() {
	flag.Parse()
	generator.Generate(*basicPath, *confFile)
}
