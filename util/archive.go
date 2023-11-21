package util

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

//func tree2string() ([]string, error) {
//	root := GetRoot()
//	var content []string
//	next := root.next()
//	for next != nil {
//		content = append(content, next.Node2String())
//		next = next.next()
//	}
//	return content, nil
//}
//
//func (node *Node) Node2String() string {
//	if node == nil {
//		return ""
//	}
//	var content string
//	if node.grade > 0 {
//		for i := 0; i < node.grade; i++ {
//			content += "#"
//		}
//		content += " "
//	}
//	content += node.content
//	return content
//}

//func (node *Node) next() *Node {
//	if node.children != nil && len(node.children) != 0 {
//		return node.children[0]
//	}
//	for {
//		rank, err := getRankofParent(node)
//		if err != nil {
//			return nil
//		}
//		if rank == -1 {
//			node = node.parent
//		} else {
//			return node.parent.children[rank+1]
//		}
//	}
//}

//// find the nth node in the tree,
//// if not found, return nil
//// n-- when recursive dfs
//func getNodeByNth(node *Node, lineNum int) (*Node, int) {
//	if node != GetRoot() {
//		lineNum--
//	}
//	if lineNum == 0 {
//		return node, lineNum
//	}
//	for _, v := range node.children {
//		var returnNode *Node
//		returnNode, lineNum = getNodeByNth(v, lineNum)
//		if lineNum == 0 {
//			return returnNode, lineNum
//		}
//	}
//	return nil, lineNum
//}

//func IsEmpty() bool {
//	return len(GetRoot().children) == 0
//}
//// return 0~len-2,-1
//func getRankofParent(node *Node) (int, error) {
//	if node.parent == nil {
//		return 0, errors.New("getRankofParent(): node.parent == nil")
//	}
//	for i, v := range node.parent.children {
//		if v == node {
//			if i == len(node.parent.children)-1 {
//				return -1, nil
//			}
//			return i, nil
//		}
//	}
//	return 0, errors.New("getRankofParent(): node.parent.children not found")
//}

//func dumpNode(node *Node) {
//	fmt.Println(node.content)
//	if len(node.children) != 0 {
//		fmt.Println("children: ")
//	}
//	for _, child := range node.children {
//		dumpNode(child)
//	}
//}
//
//func Dump() {
//	fmt.Println("Dump")
//	root := GetRoot()
//	for _, child := range root.children {
//		dumpNode(child)
//	}
//}
