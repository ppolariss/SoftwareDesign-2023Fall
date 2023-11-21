package util

import (
	. "design/interfaces"
	"errors"
	"strconv"
	"strings"
	"sync"
)

type Node struct {
	Content  string
	Children []*Node
	parent   *Node
	Grade    int
}

var root *Node
var once sync.Once

func IsInit() bool {
	if GetRoot() == nil {
		return false
	}
	//if len(GetRoot().children) == 0 {
	//	return false
	//}
	return true
}

func GetRoot() *Node {
	once.Do(func() {
		root = &Node{Content: "root", Children: []*Node{}, parent: nil, Grade: 0}
	})
	return root
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

// para content to find, node to start, nth
// return nth and node
func recurGetNodeByContent(content string, node *Node, nth int) (int, *Node) {
	if transNum(node.Content) == content {
		return nth, node
	}
	for _, child := range node.Children {
		nth++
		nth, retNode := recurGetNodeByContent(content, child, nth)
		if retNode != nil {
			return nth, retNode
		}
	}
	return 0, nil
}

// return nth and node
func GetNodeByContent(content string) (int, *Node) {
	root := GetRoot()
	for _, child := range root.Children {
		nth, retNode := recurGetNodeByContent(content, child, 1)
		if retNode != nil {
			return nth, retNode
		}
	}
	return 0, nil
}

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
	if (newNode.Grade == 0 || newNode.Grade > currentNode.Grade) && currentNode.Grade != 0 {
		newNode.parent = currentNode
		currentNode.Children = append(currentNode.Children, newNode)
		// if newNode.grade == 0 || newNode.grade == currentNode.grade+1 {
		// 	newNode.parent = currentNode
		// 	currentNode.children = append(currentNode.children, newNode)
		// } else if currentNode.grade < newNode.grade {
		// 	return errors.New("traceback(): incorrect syntax")
	} else {
		// OutputAsFile(0)

		// level up
		// due to currentNode is changed, so we need to record
		for i, times := 0, currentNode.Grade-newNode.Grade+1; i < times; i++ {
			currentNode = currentNode.parent
			if currentNode == nil {
				return errors.New("traceback(): currentNode.parent == nil")
			}
		}
		newNode.parent = currentNode
		currentNode.Children = append(currentNode.Children, newNode)
		// fmt.Println(currentNode.content)
	}
	return nil
}

func String2tree(ss []string) error {
	current := GetRoot()
	if current.Children != nil {
		current.Children = []*Node{}
	}

	for _, content := range ss {
		node, grade := ParseToNode(content)
		err := traceback(current, node)
		if err != nil {
			return errors.New(err.Error())
		}
		if grade != 0 {
			current = node
		}
	}
	return nil
}

// ParseToNode parse a line to a node,
// return the node and the number of # if valid
func ParseToNode(content string) (*Node, int) {
	grade, content := parseToString(content)
	node := Node{Content: content, Children: []*Node{}, parent: nil, Grade: grade}
	return &node, grade
}

func parseToString(content string) (int, string) {
	content = strings.TrimRight(content, "\n")
	grade := GetGrade(content)
	// remove the # and space
	if grade > 0 {
		content = content[grade+1:]
	}
	return grade, content
}

func transNum(s string) string {
	s = strings.TrimRight(s, "\n")
	ss := strings.Split(s, " ")
	if len(ss) == 1 {
		return s
	}
	// try to parse the first word to int
	_, err := strconv.Atoi(ss[0][:len(ss[0])-1])
	if err != nil {
		return s
	}
	return strings.Join(ss[1:], " ")
}

func (node *Node) Name() string {
	return node.Content
}

func (node *Node) GetChildren() []TreeOut {
	children := make([]TreeOut, len(node.Children))
	for i, child := range node.Children {
		//children = append(children, child)
		children[i] = child
	}
	return children
	//return node.children
}
