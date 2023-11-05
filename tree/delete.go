package tree

import (
	e "design/myError"
)

// Delete para line_num or content
// return nth content error
func Delete(lineNum int, content string) (int, string, error) {

	var originContent string
	if !IsInit() {
		return 0, "", e.NewMyError("delete: No file in workspace")
	}
	if lineNum > 0 {
		if lineNum > Length {
			return 0, "", e.NewMyError("delete: line number out of range")
		}
		tempNode, _ := getNodeByNth(GetRoot(), lineNum)
		originContent = tempNode.Node2String()
	} else {
		var node *Node
		lineNum, node = getNodeByContent(content)
		originContent = node.Node2String()
		if lineNum <= 0 || lineNum > Length {
			return 0, "", e.NewMyError("delete: content not match")
		}
	}

	lineNum--
	err := tree2string()
	if err != nil {
		return 0, "", err
	}
	fileContent = append(fileContent[:lineNum], fileContent[lineNum+1:]...)
	err = string2tree()
	if err != nil {
		return 0, "", err
	}
	updateLength(-1)
	return lineNum + 1, originContent, nil
}
