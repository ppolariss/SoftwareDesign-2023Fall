package util

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ReadStrings get []string from file
func ReadStrings(file *os.File) ([]string, error) {
	var result []string
	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				if content != "" {
					content = strings.TrimRight(content, "\n")
					result = append(result, content)
				}
				break
				// 注意此处还要处理最后一行
			} else {
				return result, errors.New(err.Error())
			}
		}
		if content == "" {
			continue
		}
		content = strings.TrimRight(content, "\n")
		result = append(result, content)
	}
	return result, nil
}
