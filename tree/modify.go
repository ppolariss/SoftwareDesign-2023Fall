package tree

import (
	e "design/myError"
)

func (n *Node) AddChildBefore(newNode *Node, nextNode *Node) error {
	if n == nil {
		return e.NewMyError("add_child_before: nil node")
	}
	if nextNode == nil {
		// tail
		n.children = append(n.children, newNode)
	} else {
		// insert before newNode
		for i, v := range n.children {
			if v == nextNode {
				n.children = append(n.children[:i], append([]*Node{newNode}, n.children[i:]...)...)
				break
			}
		}
	}
	return nil
}

func (n *Node) AddChildAfter(newNode *Node, preNode *Node) error {
	if n == nil {
		return e.NewMyError("add_child_after: nil node")
	}
	if preNode == nil {
		// head
		n.children = append([]*Node{newNode}, n.children...)
	} else {
		// insert after newNode
		for i, v := range n.children {
			if v == preNode {
				n.children = append(n.children[:i+1], append([]*Node{newNode}, n.children[i+1:]...)...)
				break
			}
		}
	}
	return nil
}

func Append_head(newNode *Node) error {
	if !IsInit() {
		return e.NewMyError("append_head: No file in workspace")
	}
	tree := GetRoot()
	// if IsEmpty() {
	// 	root.AddChild(newNode, nil)
	// } else {
	// 	root.AddChild(newNode, tree.children[0])
	// }
	tree.children = append([]*Node{newNode}, tree.children...)
	return nil
}

func Append_tail(newNode *Node) error {
	if !IsInit() {
		return e.NewMyError("append_tail: No file in workspace")
	}
	end := GetRoot()
	for {
		if end.children == nil || len(end.children) == 0 {
			break
		} else {
			end = end.children[len(end.children)-1]
		}
	}
	traceback(end, newNode)
	return nil
}

// func (n *Node) DeleteChild(child *Node) {
// 	for i, v := range n.children {
// 		if v == child {
// 			n.children = append(n.children[:i], n.children[i+1:]...)
// 			break
// 		}
// 	}
// }

// find the nth node in the tree,
// if not found, return nil
// n-- when recursive dfs
func dfs(node *Node, line_num int) (*Node, int) {
	line_num--
	if line_num == 0 {
		return node, line_num
	}
	for _, v := range node.children {
		var return_node *Node
		return_node, line_num = dfs(v, line_num)
		if line_num == 0 {
			return return_node, line_num
		}
	}
	return nil, line_num
}

func Insert(line_num int, newNode *Node) error {
	if line_num > Length+1 || line_num < 1 {
		return e.NewMyError("insert: line number out of range")
	}
	if !IsInit() {
		return e.NewMyError("insert: No file in workspace")
	}

	// insert the nth line, get the n-1th node
	// the root node is not counted
	prenode, _ := dfs(GetRoot(), line_num)
	if prenode == nil {
		return e.NewMyError("insert: no such node")
	}

	if newNode.grade == 0 {
		prenode.children = append([]*Node{newNode}, prenode.children...)
		newNode.parent = prenode
		return nil
	}

	if prenode.grade == 0 && prenode != GetRoot() {
		if newNode.grade != prenode.parent.grade+1 {
			return e.NewMyError("insert: incorrect syntax")
		}
		prenode.parent.AddChildAfter(newNode, prenode)
		newNode.parent = prenode.parent
		return nil
	}
	
	// head
	if newNode.grade == prenode.grade+1 {
		// adopt the continuous 0-grade child
		for _, v := range prenode.children {
			if v.grade == 0 {
				adoptChild(newNode, v)
			} else {
				break
			}
		}
		prenode.children = append([]*Node{newNode}, prenode.children...)
		newNode.parent = prenode
	} else if newNode.grade > prenode.grade+1 {
		return e.NewMyError("insert: incorrect syntax")
	} else if prenode.grade == newNode.grade {
		// adopt all the children
		stealChildren(newNode, prenode)
		prenode.parent.AddChildAfter(newNode, prenode)
		newNode.parent = prenode.parent
	} else {
		// valid only when the children of prenode are all 0-grade
		for _, v := range prenode.children {
			if v.grade != 0 {
				return e.NewMyError("insert: incorrect syntax")
			}
		}
		stealChildren(newNode, prenode)
		for {
			parent := prenode.parent
			if parent == nil {
				return e.NewMyError("insert: incorrect syntax")
			}
			if parent.grade < newNode.grade {
				parent.AddChildAfter(newNode, prenode)
				newNode.parent = parent
				break
			}
		}
		// 重排后续所有节点
	}
	return nil
	// if line_num == 0 {
	// 	tree.AddChild(newNode, tree.children[0])
	// } else {
	// 	for i, v := range tree.children {
	// 		if i == line_num-1 {
	// 			tree.AddChild(newNode, v)
	// 			break
	// 		}
	// 	}
	// }
	// return nil
}
