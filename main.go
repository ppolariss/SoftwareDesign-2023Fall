package main

import (
	"design/command"
	"fmt"
	"os"
	//"os"
)

func init() {
	dirPaths := []string{
		"./files",
		"./testFiles",
		"./logFiles",
	}
	for _, dirPath := range dirPaths {
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// 目录不存在，创建目录,权限设置
			_ = os.Mkdir(dirPath, 0755)
		}
	}
}

func main() {
	err := command.Do()
	if err != nil {
		fmt.Println(err)
	}
}
