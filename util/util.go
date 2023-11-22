package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func Output(content string, file *os.File) error {
	if file == nil {
		file = os.Stdout
	}
	_, err := file.WriteString(content)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

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

func GetInterval(nowTime string, createTime string) (string, error) {
	nowTime = strings.TrimRight(nowTime, "\n")
	createTime = strings.TrimRight(createTime, "\n")
	layout := "20060102 15:04:05"
	now, err := time.Parse(layout, nowTime)
	if err != nil {
		return "", errors.New(err.Error())
	}
	create, err := time.Parse(layout, createTime)
	if err != nil {
		return "", errors.New(err.Error())
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

// OutputAsFile para: 0: output to terminal; 1: output to file
func AsFile(para int, fileContent []string, filePath string) error {
	var file *os.File
	var err error
	if para == 1 {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return errors.New(err.Error())
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		// 清空文件内容
		err = file.Truncate(0)
		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		file = os.Stdout
	}
	for _, line := range fileContent {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return errors.New(err.Error())
		}
	}

	//err := util.String2tree(fileContent)
	//if err != nil {
	//	return err
	//}
	//if !util.IsInit() {
	//	return errors.New("OutputAsFile(): No file in workspace")
	//}
	//
	//tree := util.GetRoot()
	//if para != 0 {
	//	// 向文件输入
	//
	//	for _, child := range tree.Children {
	//		recurOutputAsFile(child, file)
	//	}
	//
	//} else {
	//	// 向终端输入
	//	for _, child := range tree.Children {
	//		recurOutputAsFile(child, nil)
	//	}
	//}
	return nil
}

//func recurOutputAsFile(node *util.Node, file *os.File) {
//	if node.Grade != 0 {
//		for i := 0; i < node.Grade; i++ {
//			_ = util.Output("#", file)
//
//		}
//		_ = util.Output(" ", file)
//
//	}
//
//	_ = util.Output(node.Content+"\n", file)
//	for _, child := range node.Children {
//		recurOutputAsFile(child, file)
//	}
//}
