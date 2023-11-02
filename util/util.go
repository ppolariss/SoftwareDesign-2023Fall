package util

import (
	// "bufio"
	"bufio"
	e "design/myError"
	"fmt"
	"os"
	"strings"
	"time"
	// "github.com/go-delve/delve/pkg/dwarf/reader"
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

	// return Unix2Time(interval), nil
}

const formatStr = "20060102 15:04:05"

func GetNow() string {
	return time.Now().Format(formatStr)
}

// func tree() {
// 	// ├── 新的标题
// 	// └── 我的书签
// 	// 	└── 学习资源
// 	// 		├── 新的⼦标题
// 	// 		└── 编程
// 	// 			└── ·新的⽂本
// }

// func main() {
// 	write()

// 	var path string = "test.txt"
// 	file, err := os.Open(path)
// 	defer file.Close()
// 	if err != nil {
// 		file, _ = os.Create(path)
// 	}
// 	// bufio ioutil
// 	reader := bufio.NewReader(file)
// 	content, err := reader.ReadString('\n')
// 	if err != nil {
// 		// panic(err)
// 		fmt.Println(err)
// 	}
// 	fmt.Println(content)

// }

// // 时间转化
// // 树结构输出

// func write() {
// 	file_path := "test.txt"

// 	file, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 		// log.Fatal(err)
// 	}
// 	defer file.Close()

// 	if _, err := file.WriteString("This is some extra text."); err != nil {
// 		panic(err)
// 	}
// }
