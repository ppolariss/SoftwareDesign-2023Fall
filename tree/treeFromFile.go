package tree

import (
	"bufio"
	e "design/myError"
	"os"
)

// when adding a new node, the position of the node is determined by the current node and the grade,
// only used when append
func traceback(currentNode *Node, newNode *Node) error {
	// if currentNode.grade == 0 {
	// 	for {
	// 		currentNode = currentNode.parent
	// 		if currentNode == nil {
	// 			return e.NewMyError("traceback(): currentNode.parent == nil")
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
	if currentNode == nil || newNode == nil {
		return e.NewMyError("traceback(): currentNode == nil || newNode == nil")
	}
	// level down
	if (newNode.grade == 0 || newNode.grade > currentNode.grade) && currentNode.grade != 0 {
		newNode.parent = currentNode
		currentNode.children = append(currentNode.children, newNode)
		// if newNode.grade == 0 || newNode.grade == currentNode.grade+1 {
		// 	newNode.parent = currentNode
		// 	currentNode.children = append(currentNode.children, newNode)
		// } else if currentNode.grade < newNode.grade {
		// 	return e.NewMyError("traceback(): incorrect syntax")
	} else {
		// OutputAsFile(0)

		// level up
		// due to currentNode is changed, so we need to record
		for i, times := 0, currentNode.grade-newNode.grade+1; i < times; i++ {
			currentNode = currentNode.parent
			if currentNode == nil {
				return e.NewMyError("traceback(): currentNode.parent == nil")
			}
		}
		newNode.parent = currentNode
		currentNode.children = append(currentNode.children, newNode)
		// fmt.Println(currentNode.content)
	}
	return nil
}

// parse a new file to a tree
// return the length and error
func parseFromFile(file_path string) (int, error) {
	root := GetRoot()
	if root.children != nil {
		root.children = []*Node{}
	}

	file, err := os.OpenFile(file_path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, e.NewMyError(err.Error())
	}
	defer file.Close()

	current := root
	reader := bufio.NewReader(file)
	return_flag := false
	count := 0
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				if content == "" {
					break
				}
				return_flag = true
				// 注意此处还要处理最后一行
			} else {
				return count, e.NewMyError(err.Error())
			}
		}
		// if content == "" {
		// 	continue
		// }
		count++

		fileContent = append(fileContent, content)
		node, grade, err := ParseNode(content)
		if err != nil {
			return count, e.NewMyError(err.Error())
		}

		err = traceback(current, node)
		if err != nil {
			return count, e.NewMyError(err.Error())
		}

		if grade != 0 {
			current = node
		}
		if return_flag {
			break
		}
	}
	// Dump()
	return count, nil
}
