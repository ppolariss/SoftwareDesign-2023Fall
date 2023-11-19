package output

import (
	"errors"
	// "fmt"
	"strconv"
	"strings"
)

func string2tree(ss []string) error {
	current := GetRoot()
	if current.children != nil {
		current.children = []*Node{}
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

func tree2string() ([]string, error) {
	root := GetRoot()
	var content []string
	next := root.next()
	for next != nil {
		content = append(content, next.Node2String())
		next = next.next()
	}
	return content, nil
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

// ParseToNode parse a line to a node,
// return the node and the number of # if valid
func ParseToNode(content string) (*Node, int) {
	grade, content := parseToString(content)
	node := Node{content: content, children: []*Node{}, parent: nil, grade: grade}
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
