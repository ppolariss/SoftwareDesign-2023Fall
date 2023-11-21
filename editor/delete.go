package editor

import (
	"design/util"
	"errors"
)

// Delete para line_num or content
// return index content error
func Delete(lineNum int, content string, fileContent *[]string) (int, string, error) {

	var originContent string
	if !util.IsInit() {
		return 0, "", errors.New("delete: No file in workspace")
	}
	if lineNum > 0 {
		if lineNum > len(*fileContent) {
			return 0, "", errors.New("delete: line number out of range")
		}
		originContent = (*fileContent)[lineNum-1]
	} else {
		var err error
		lineNum, err = matchContent(content, *fileContent)
		if err != nil {
			return 0, "", err
		}
		if lineNum <= 0 || lineNum > len(*fileContent) {
			return 0, "", errors.New("delete: content not match")
		}
		originContent = (*fileContent)[lineNum-1]
	}
	*fileContent = append((*fileContent)[:lineNum-1], (*fileContent)[lineNum:]...)
	//err = string2tree()
	//if err != nil {
	//	return 0, "", err
	//}
	return lineNum, originContent, nil
}
