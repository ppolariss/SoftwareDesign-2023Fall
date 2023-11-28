package main

import (
	"design/command"
	"fmt"
	"os"
)

func init() {
	dirPaths := []string{
		"./files",
		"./testFiles",
		"./logFiles",
		"./logFiles/log",
	}
	for _, dirPath := range dirPaths {
		if fileInfo, err := os.Stat(dirPath); os.IsNotExist(err) || !fileInfo.IsDir() {
			// 目录不存在，创建目录,权限设置
			_ = os.Mkdir(dirPath, 0755)
		}
	}

	command.Init()
}

func main() {
	err := command.Do()
	if err != nil {
		fmt.Println(err)
	}
}
