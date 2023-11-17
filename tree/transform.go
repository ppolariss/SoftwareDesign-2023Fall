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
		node, grade := ParseToNode(content)
		err := traceback(current, node)
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

func getBareContent(s string) string {
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return s
	}
	i := 0
	for {
		if s[i] == '#' {
			i++
		} else if s[i] == ' ' {
			break
		} else {
			i = 0
			break
		}
	}
	if i > 0 {
		s = s[i+1:]
	}

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
