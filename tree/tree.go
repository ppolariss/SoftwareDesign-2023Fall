package tree

import (
	// "path/filepath"
	e "design/myError"
	"strings"
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

//  = &Node{content: "root", children: []*Node{}, parent: nil}

func Load(path string) error {
	file_path = path
	var err error
	Length, err = parseFromFile(path)
	return err
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

// parse a line to a node,
// return the node and the number of # if valid
func ParseNode(content string) (*Node, int, error) {
	content = strings.TrimRight(content, "\n")
	grade := GetGrade(content)
	// remove the # and space
	if grade > 0 {
		content = content[grade+1:]
	}
	node := Node{content: content, children: []*Node{}, parent: nil, grade: grade}
	return &node, grade, nil
}

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

func recurGetNodeByContent(content string, node *Node) *Node {
	if node.content == content {
		return node
	}
	for _, child := range node.children {
		retNode := recurGetNodeByContent(content, child)
		if retNode != nil {
			return retNode
		}
	}
	return nil
}

func getNodeByContent(content string) (*Node) {
	root := GetRoot()
	for _, child := range root.children {
		retNode := recurGetNodeByContent(content, child)
		if retNode != nil {
			return retNode
		}
	}
	return nil
}