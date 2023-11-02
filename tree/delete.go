package tree

import (
	e "design/myError"
)

func removeFromParent(node *Node) {
	if node.parent == nil {
		return
	}
	for i, v := range node.parent.children {
		if v == node {
			node.parent.children = append(node.parent.children[:i], node.parent.children[i+1:]...)
			break
		}
	}
	node.parent = nil
}

func adoptChild(node *Node, child *Node) {
	removeFromParent(child)
	node.children = append(node.children, child)
	child.parent = node
}

func stealChildren(node *Node, parent *Node) {
	node.children = append(node.children, parent.children...)
	for _, child := range node.children {
		child.parent = node
	}
	parent.children = []*Node{}
}

// para line_num or content
// return nth content error
func Delete(line_num int, content string) (int, string, error) {

	var originContent string
	if !IsInit() {
		return 0, "", e.NewMyError("delete: No file in workspace")
	}
	if line_num > 0 {
		if line_num > Length {
			return 0, "", e.NewMyError("delete: line number out of range")
		}
		temp_node, _ := getNodeByNth(GetRoot(), line_num)
		originContent = temp_node.Node2String()
	} else {
		var node *Node
		line_num, node = getNodeByContent(content)
		originContent = node.Node2String()
		if line_num <= 0 || line_num > Length {
			return 0, "", e.NewMyError("delete: content not match")
		}
	}

	line_num--
	err := tree2string()
	if err != nil {
		return 0, "", err
	}
	fileContent = append(fileContent[:line_num], fileContent[line_num+1:]...)
	err = string2tree()
	if err != nil {
		return 0, "", err
	}
	updateLength(-1)
	return line_num + 1, originContent, nil
}
