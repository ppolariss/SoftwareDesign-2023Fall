package tree

import (
	// "path/filepath"
	e "design/myError"
	"sync"
)

// 要不要改大写
type Node struct {
	content  string
	children []*Node
	parent   *Node
	grade    int
}

var file_path string
var Length int
var root *Node
var once sync.Once
var fileContent []string

//  = &Node{content: "root", children: []*Node{}, parent: nil}

func Load(path string) error {
	file_path = path
	var err error
	Length, err = parseFromFile(path)
	return err
}

func updateLength(num int) int {
	if num != 0 {
		Length += num
	} else {
		Length = len(fileContent)
	}
	return Length
}

func IsInit() bool {
	if file_path == "" {
		return false
	} else {
		return true
	}
}

func GetRoot() *Node {
	once.Do(func() {
		root = &Node{content: "root", children: []*Node{}, parent: nil, grade: 0}
	})
	return root
}

func IsEmpty() bool {
	if GetRoot().children == nil || len(GetRoot().children) == 0 {
		return true
	} else {
		return false
	}
}

// return the number of # if valid
func GetGrade(content string) int {
	i := 0
	for {
		if content[i] == '#' {
			i++
		} else if content[i] == ' ' {
			return i
		} else {
			return 0
		}
	}
}

// return 0~len-2,-1
func getRankofParent(node *Node) (int, error) {
	if node.parent == nil {
		return 0, e.NewMyError("getRankofParent(): node.parent == nil")
	}
	for i, v := range node.parent.children {
		if v == node {
			if i == len(node.parent.children)-1 {
				return -1, nil
			}
			return i, nil
		}
	}
	return 0, e.NewMyError("getRankofParent(): node.parent.children not found")
}

// para content to find, node to start, nth
// return nth and node
func recurGetNodeByContent(content string, node *Node, nth int) (int, *Node) {
	if transNum(node.content) == content {
		return nth, node
	}
	for _, child := range node.children {
		nth++
		nth, retNode := recurGetNodeByContent(content, child, nth)
		if retNode != nil {
			return nth, retNode
		}
	}
	return 0, nil
}

// return nth and node
func getNodeByContent(content string) (int, *Node) {
	root := GetRoot()
	for _, child := range root.children {
		nth, retNode := recurGetNodeByContent(content, child, 1)
		if retNode != nil {
			return nth, retNode
		}
	}
	return 0, nil
}

func (node *Node) next() *Node {
	if node.children != nil && len(node.children) != 0 {
		return node.children[0]
	}
	for {
		rank, err := getRankofParent(node)
		if err != nil {
			return nil
		}
		if rank == -1 {
			node = node.parent
		} else {
			return node.parent.children[rank+1]
		}
	}
}
