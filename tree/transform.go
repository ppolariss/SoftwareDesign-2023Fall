package tree

import (
	e "design/myError"
	// "fmt"
	"strconv"
	"strings"
)

func string2tree() error {
	current := GetRoot()
	if current.children != nil {
		current.children = []*Node{}
	}

	for _, content := range fileContent {
		node, grade, err := ParseNode(content)
		if err != nil {
			return e.NewMyError(err.Error())
		}
		err = traceback(current, node)
		if err != nil {
			return e.NewMyError(err.Error())
		}
		if grade != 0 {
			current = node
		}
	}
	return nil
}

func tree2string() error {
	root := GetRoot()
	fileContent = []string{}
	next := root.next()
	for next != nil {
		fileContent = append(fileContent, next.Node2String())
		next = next.next()
	}
	return nil
}

func (node *Node) Node2String() string {
	if node == nil {
		return ""
	}
	var content string
	if node.grade > 0 {
		for i := 0; i < node.grade; i++ {
			content += "#"
		}
		content += " "
	}
	content += node.content
	return content
}

// ParseNode parse a line to a node,
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

func transNum(s string) string {
	ss := strings.Split(s, " ")
	if len(ss) == 1 {
		return s
	}
	_, err := strconv.Atoi(ss[0][:len(ss[0])-1])
	if err != nil {
		return s
	}
	return strings.Join(ss[1:], " ")
}
