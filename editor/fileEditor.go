package editor

import (
	"errors"
	"strconv"
	"strings"
)

//var fileContent []string

//var filePath string
//
//func Load(path string) error {
//	filePath = path
//	fileContent = fileContent[:0]
//	// may drop the empty line
//	fileBytes, err := os.ReadFile(filePath)
//	if err != nil {
//		return err
//	}
//	lines := bytes.Split(fileBytes, []byte("\n"))
//	for _, line := range lines {
//		if len(line) == 0 {
//			continue
//		}
//		fileContent = append(fileContent, string(line))
//	}
//	//file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
//	//if err != nil {
//	//	return errors.New(err.Error())
//	//}
//	//defer func(file *os.File) {
//	//	_ = file.Close()
//	//}(file)
//	//
//	//reader := bufio.NewReader(file)
//	//for {
//	//	content, err := reader.ReadString('\n')
//	//	if err != nil {
//	//		if err.Error() != "EOF" {
//	//			return errors.New(err.Error())
//	//		}
//	//
//	//		if content == "" {
//	//			break
//	//		} else {
//	//			fileContent = append(fileContent, content)
//	//			break
//	//		}
//	//		// 注意此处还要处理最后一行
//	//	}
//	//	// if content == "" {
//	//	// 	continue
//	//	// }
//	//	fileContent = append(fileContent, content)
//	//}
//	return nil
//}
//
//func IsInit() bool {
//	return filePath != ""
//}

//func getLength() int {
//	return len(fileContent)
//}

// return lineNum according to content
func matchContent(content string, fileContent []string) (int, error) {
	for i, v := range fileContent {
		s := getBareContent(v)
		if s == content {
			return i + 1, nil
		}
	}
	return 0, errors.New("matchContent(): content not found")
}

func getBareContent(s string) string {
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return s
	}
	i := 0
	for {
		if s[i] == '#' {
			i++
		} else if s[i] == ' ' {
			break
		} else {
			i = 0
			break
		}
	}
	if i > 0 {
		s = s[i+1:]
	}

	ss := strings.Split(s, " ")
	if len(ss) == 1 {
		return s
	}
	// try to parse the first word to int
	_, err := strconv.Atoi(ss[0][:len(ss[0])-1])
	if err != nil {
		return s
	}
	return strings.Join(ss[1:], " ")
}

//func UpdateFileContent(newFileContent []string) {
//	fileContent = newFileContent
//}
//
//func GetFileContent() []string {
//	return fileContent
//}
