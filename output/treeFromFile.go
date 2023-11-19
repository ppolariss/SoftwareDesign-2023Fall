package output

import (
	"errors"
)

// when adding a new node, the position of the node is determined by the current node and the grade,
// only used when append
func traceback(currentNode *Node, newNode *Node) error {
	// if currentNode.grade == 0 {
	// 	for {
	// 		currentNode = currentNode.parent
	// 		if currentNode == nil {
	// 			return errors.New("traceback(): currentNode.parent == nil")
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
	if currentNode == nil || newNode == nil {
		return errors.New("traceback(): currentNode == nil || newNode == nil")
	}
	// level down
	if (newNode.grade == 0 || newNode.grade > currentNode.grade) && currentNode.grade != 0 {
		newNode.parent = currentNode
		currentNode.children = append(currentNode.children, newNode)
		// if newNode.grade == 0 || newNode.grade == currentNode.grade+1 {
		// 	newNode.parent = currentNode
		// 	currentNode.children = append(currentNode.children, newNode)
		// } else if currentNode.grade < newNode.grade {
		// 	return errors.New("traceback(): incorrect syntax")
	} else {
		// OutputAsFile(0)

		// level up
		// due to currentNode is changed, so we need to record
		for i, times := 0, currentNode.grade-newNode.grade+1; i < times; i++ {
			currentNode = currentNode.parent
			if currentNode == nil {
				return errors.New("traceback(): currentNode.parent == nil")
			}
		}
		newNode.parent = currentNode
		currentNode.children = append(currentNode.children, newNode)
		// fmt.Println(currentNode.content)
	}
	return nil
}

//// parse a new file to a tree
//// return the length and error
//func parseFromFile(filePath string) error {
//	root := GetRoot()
//	if root.children != nil {
//		root.children = []*Node{}
//	}
//
//
//
//	current := root
//
//	count := 0
//	for {
//		content, err := reader.ReadString('\n')
//		if err != nil {
//			if err.Error() == "EOF" {
//				if content == "" {
//					break
//				}
//				returnFlag = true
//				// 注意此处还要处理最后一行
//			} else {
//				return errors.New(err.Error())
//			}
//		}
//		// if content == "" {
//		// 	continue
//		// }
//		count++
//
//		fileEditor.fileContent = append(fileEditor.fileContent, content)
//		node, grade := ParseToNode(content)
//		if err != nil {
//			return errors.New(err.Error())
//		}
//
//		err = traceback(current, node)
//		if err != nil {
//			return errors.New(err.Error())
//		}
//
//		if grade != 0 {
//			current = node
//		}
//		if returnFlag {
//			break
//		}
//	}
//	// Dump()
//	// below is written at 2023/11/18
//	err = tree2string()
//	if err != nil {
//		return err
//	}
//	return nil
//}
