package tree

import (
	e "design/myError"
	// "slices"
	// "fmt"
)

func AppendHead(newNode *Node) (int, error) {
	if !IsInit() {
		return 0, e.NewMyError("append_head: No file in workspace")
	}
	tree := GetRoot()
	if newNode.grade != 0 && newNode.grade != 1 {
		return 0, e.NewMyError("append_head: incorrect syntax")
	}
	// if IsEmpty() {
	// 	root.AddChild(newNode, nil)
	// } else {
	// 	root.AddChild(newNode, tree.children[0])
	// }
	tree.children = append([]*Node{newNode}, tree.children...)
	newNode.parent = tree
	updateLength(1)
	return 1, nil
}

func AppendTail(newNode *Node) (int, error) {
	if !IsInit() {
		return 0, e.NewMyError("append_tail: No file in workspace")
	}
	end := GetRoot()
	for {
		if end.children == nil || len(end.children) == 0 {
			break
		} else {
			end = end.children[len(end.children)-1]
			if end.grade == 0 {
				end = end.parent
				break
			}
		}
	}
	err := traceback(end, newNode)
	if err != nil {
		return 0, err
	}
	// tree2string()
	// string2tree()
	return updateLength(1), nil
}

// find the nth node in the tree,
// if not found, return nil
// n-- when recursive dfs
func getNodeByNth(node *Node, lineNum int) (*Node, int) {
	if node != GetRoot() {
		lineNum--
	}
	if lineNum == 0 {
		return node, lineNum
	}
	for _, v := range node.children {
		var returnNode *Node
		returnNode, lineNum = getNodeByNth(v, lineNum)
		if lineNum == 0 {
			return returnNode, lineNum
		}
	}
	return nil, lineNum
}

func Insert(lineNum int, newNode *Node) (int, error) {

	if lineNum > Length+1 || lineNum < 1 {
		return 0, e.NewMyError("insert: line number out of range")
	}
	if !IsInit() {
		return 0, e.NewMyError("insert: No file in workspace")
	}

	err := tree2string()
	if err != nil {
		return 0, e.NewMyError(err.Error())
	}

	// n-1 node before newline
	// get the 0~n-2 lines
	// fileContent = slices.Insert(fileContent, lineNum, newNode.Node2String())
	fileContent = append(fileContent[:lineNum-1], append([]string{newNode.Node2String()}, fileContent[lineNum-1:]...)...)

	err = string2tree()
	if err != nil {
		return 0, err
	}
	updateLength(1)
	return lineNum, nil
}
