package editor

import (
	"errors"

	"design/output"
)

func Insert(lineNum int, content string) (int, error) {
	if !output.IsInit() {
		return 0, errors.New("insert: No file in workspace")
	}
	if lineNum == -1 {
		fileContent = append(fileContent, content)
		return getLength(), nil
	}
	if lineNum > getLength()+1 || lineNum < 1 {
		return 0, errors.New("insert: line number out of range")
	}

	// n-1 node before newline
	// get the 0~n-2 lines
	// fileContent = slices.Insert(fileContent, lineNum, newNode.Node2String())
	fileContent = append(fileContent[:lineNum-1], append([]string{content}, fileContent[lineNum-1:]...)...)

	//err = string2tree()
	//if err != nil {
	//	return err
	//}
	return lineNum, nil
}