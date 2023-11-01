package tree

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
