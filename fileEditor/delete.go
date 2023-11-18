package fileEditor

import (
	e "design/myError"
	"design/output"
)

// Delete para line_num or content
// return index content error
func Delete(lineNum int, content string) (int, string, error) {

	var originContent string
	if !output.IsInit() {
		return 0, "", e.NewMyError("delete: No file in workspace")
	}
	if lineNum > 0 {
		if lineNum > getLength() {
			return 0, "", e.NewMyError("delete: line number out of range")
		}
		originContent = fileContent[lineNum-1]
	} else {
		var err error
		lineNum, err = matchContent(content)
		if err != nil {
			return 0, "", err
		}
		if lineNum <= 0 || lineNum > getLength() {
			return 0, "", e.NewMyError("delete: content not match")
		}
		originContent = fileContent[lineNum-1]
	}
	fileContent = append(fileContent[:lineNum-1], fileContent[lineNum:]...)
	//err = string2tree()
	//if err != nil {
	//	return 0, "", err
	//}
	return lineNum, originContent, nil
}
