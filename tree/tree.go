package tree

import (
	e "design/myError"
	"sync"
)

type Node struct {
	content  string
	children []*Node
	parent   *Node
	grade    int
}

var filePath string

// var Length int
var root *Node
var once sync.Once
var fileContent []string

func Load(path string) error {
	filePath = path
	var err error
	_, err = parseFromFile(path)
	return err
}

//func updateLength(num int) int {
//	if num != 0 {
//		Length += num
//	} else {
//		Length = len(fileContent)
//	}
//	return Length
//}

func IsInit() bool {
	return filePath != ""
}

func GetRoot() *Node {
	once.Do(func() {
		root = &Node{content: "root", children: []*Node{}, parent: nil, grade: 0}
	})
	return root
}

func IsEmpty() bool {
	return len(GetRoot().children) == 0
}

// GetGrade return the number of # if valid
func GetGrade(content string) int {
	if content == "" {
		return 0
	}
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

// return lineNum according to content
func matchContent(content string) (int, error) {
	for i, v := range fileContent {
		s := getBareContent(v)
		if s == content {
			return i + 1, nil
		}
	}
	return 0, e.NewMyError("matchContent(): content not found")
}

func getLength() int {
	return len(fileContent)
}
