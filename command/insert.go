package command

import (
	e "design/myError"
	"design/tree"
	"fmt"
	"strconv"
	"strings"
)

type insert struct {
	// -1 表示最后
	line_num int
	content  string
}

func (c *insert) Execute() error {
	fmt.Println("insert")
	// return nil
	if c.line_num == -1 {
		n, _, err := tree.ParseNode(c.content)
		if err != nil {
			return err
		}
		return tree.Append_tail(n)
	} else {
		n, _, err := tree.ParseNode(c.content)
		if err != nil {
			return err
		}
		return tree.Insert(c.line_num, n)
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
		}
		// The line number is specified
		// 要得到一个最大整数
		if num < 0 || num > 10000 {
			return e.NewMyError("insert: line number error")
		}
		c.line_num = num
		slice_args := args[2:]
		c.content = strings.Join(slice_args, " ")
	}
	return nil
}


type append_head struct {
	content string
}

func (c *append_head) Execute() error {
	// 是否会破坏文本结构
	// 如果破坏，在哪里报错
	fmt.Println("append_head")
	n, _, err := tree.ParseNode(c.content)
	if err != nil {
		return err
	}
	return tree.Append_head(n)
	// return nil
}

func (c *append_head) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("append_head: args error")
	}
	slice_args := args[1:]
	c.content = strings.Join(slice_args, " ")
	return nil
}


type append_tail struct {
	content string
}

func (c *append_tail) Execute() error {
	fmt.Println("append_tail")
	n, _, err := tree.ParseNode(c.content)
	if err != nil {
		return err
	}
	return tree.Append_tail(n)
}

func (c *append_tail) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("append_tail: args error")
	}
	slice_args := args[1:]
	c.content = strings.Join(slice_args, " ")
	return nil
}
