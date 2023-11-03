package command

import (
	e "design/myError"
	"design/tree"
	// "fmt"
	"strconv"
	"strings"
)

type insert struct {
	// -1 表示最后
	line_num int
	content  string
}

func (c *insert) Execute() (Command, error) {
	if c.line_num == -1 {
		n, _, err := tree.ParseNode(c.content)
		if err != nil {
			return nil, err
		}
		num, err := tree.Append_tail(n)
		if err != nil {
			return nil, err
		}
		delete := &delete{line_num: num}
		return delete, err
	} else {
		n, _, err := tree.ParseNode(c.content)
		if err != nil {
			return nil, err
		}
		// println(c.line_num)
		num, err := tree.Insert(c.line_num, n)
		if err != nil {
			return nil, err
		}
		delete := &delete{line_num: num}
		return delete, err
	}
}

func (c *insert) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("insert: args error")
	} else if len(args) == 2 {
		c.line_num = -1
		c.content = args[1]
	} else {
		num, err := strconv.Atoi(args[1])
		// The line number is not specified
		if err != nil {
			c.line_num = -1
			slice_args := args[1:]
			c.content = strings.Join(slice_args, " ")
			return nil
		}
		// The line number is specified
		// 要得到一个最大整数
		if num <= 0 || num > tree.Length+1 {
			return e.NewMyError("insert: line number error")
		}
		c.line_num = num
		slice_args := args[2:]
		c.content = strings.Join(slice_args, " ")
	}
	return nil
}

func (c *insert) CallSelf() string {
	retStr := "insert"
	if c.line_num > 0 {
		retStr += " " + strconv.Itoa(c.line_num)
	}
	retStr += " " + c.content
	return retStr
}

type append_head struct {
	content string
}

func (c *append_head) Execute() (Command, error) {
	// 是否会破坏文本结构
	// 如果破坏，在哪里报错
	n, _, err := tree.ParseNode(c.content)
	if err != nil {
		return nil, err
	}
	num, err := tree.Append_head(n)
	if err != nil {
		return nil, err
	}
	delete := &delete{line_num: num}
	return delete, err
}

func (c *append_head) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("append_head: args error")
	}
	slice_args := args[1:]
	c.content = strings.Join(slice_args, " ")
	return nil
}

func (c *append_head) CallSelf() string {
	retStr := "append-head"
	retStr += " " + c.content
	return retStr
}

type append_tail struct {
	content string
}

func (c *append_tail) Execute() (Command, error) {
	n, _, err := tree.ParseNode(c.content)
	if err != nil {
		return nil, err
	}
	num, err := tree.Append_tail(n)
	if err != nil {
		return nil, err
	}
	delete := &delete{line_num: num}
	return delete, err
}

func (c *append_tail) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("append_tail: args error")
	}
	slice_args := args[1:]
	c.content = strings.Join(slice_args, " ")
	return nil
}

func (c *append_tail) CallSelf() string {
	retStr := "append-tail"
	retStr += " " + c.content
	return retStr
}
