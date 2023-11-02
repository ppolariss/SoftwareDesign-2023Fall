package command

import (
	e "design/myError"
	"design/tree"
	// "fmt"
	"strconv"
	"strings"
	
)

type delete struct {
	line_num int
	content  string //node.content
}

func (c *delete) Execute() (Command, error) {
	nth, content, err := tree.Delete(c.line_num, c.content)
	if err != nil {
		return nil, err
	}
	insert := &insert{line_num: nth, content: content}
	return insert, nil
	// 删除指定标题或⽂本。如果指定⾏号，则删除指定⾏。当删除的是标题时，其⼦标题
	// 和内容不会被删除。
}

func (c *delete) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("delete: args error")
	}

	// parse to line_num if and only if the rest is a positive number
	if len(args) == 2 {
		num, err := strconv.Atoi(args[1])
		if err == nil && num > 0 {
			if num > tree.Length {
				return e.NewMyError("delete: line number error")
			}
			c.line_num = num
			c.content = ""
			return nil
		}
	}

	// The line number is not specified
	c.line_num = 0
	slice_args := args[1:]
	c.content = strings.Join(slice_args, " ")

	return nil
}

func (c *delete) CallSelf() string {
	retStr := "delete"
	if c.line_num > 0  {
		retStr += " " + strconv.Itoa(c.line_num)
	} else {
		retStr += " " + c.content
	}
	return retStr
}
