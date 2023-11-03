package util

import (
	"bufio"
	e "design/myError"
	"fmt"
	"os"
	"strings"
	"time"
)

func Output(content string, file *os.File) error {
	if file == nil {
		fmt.Print(content)
	} else {
		_, err := file.WriteString(content)
		if err != nil {
			return e.NewMyError(err.Error())
		}
	}
	return nil
}

// get []string from file
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
				return result, e.NewMyError(err.Error())
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

func GetInterval(nowTime string, createTime string) (string, error) {
	nowTime = strings.TrimRight(nowTime, "\n")
	createTime = strings.TrimRight(createTime, "\n")
	layout := "20060102 15:04:05"
	now, err := time.Parse(layout, nowTime)
	if err != nil {
		return "", e.NewMyError(err.Error())
	}
	create, err := time.Parse(layout, createTime)
	if err != nil {
		return "", e.NewMyError(err.Error())
	}

	duration := now.Sub(create)
	hours := int(duration.Hours())
	days := hours / 24
	hours = hours % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	retStr := ""
	if days != 0 {
		retStr += fmt.Sprintf("%d天", days)
	}
	if hours != 0 {
		retStr += fmt.Sprintf("%d小时", hours)
	}
	if minutes != 0 {
		retStr += fmt.Sprintf("%d分钟", minutes)
	}
	if seconds != 0 {
		retStr += fmt.Sprintf("%d秒", seconds)
	}
	if retStr == "" {
		retStr = "0秒"
	}
	return retStr, nil
}

const formatStr = "20060102 15:04:05"

func GetNow() string {
	return time.Now().Format(formatStr)
}
