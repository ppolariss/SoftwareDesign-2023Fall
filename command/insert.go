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
	lineNum int
	content string
}

func (c *insert) Execute() (Command, error) {
	if c.lineNum == -1 {
		n, _, err := tree.ParseNode(c.content)
		if err != nil {
			return nil, err
		}
		num, err := tree.AppendTail(n)
		if err != nil {
			return nil, err
		}
		deleteC := &deleteCommand{lineNum: num}
		return deleteC, err
	} else {
		n, _, err := tree.ParseNode(c.content)
		if err != nil {
			return nil, err
		}
		// println(c.line_num)
		num, err := tree.Insert(c.lineNum, n)
		if err != nil {
			return nil, err
		}
		deleteC := &deleteCommand{lineNum: num}
		return deleteC, err
	}
}

func (c *insert) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("insert: args error")
	} else if len(args) == 2 {
		c.lineNum = -1
		c.content = args[1]
	} else {
		num, err := strconv.Atoi(args[1])
		// The line number is not specified
		if err != nil {
			c.lineNum = -1
			sliceArgs := args[1:]
			c.content = strings.Join(sliceArgs, " ")
			return nil
		}
		// The line number is specified
		// 要得到一个最大整数
		if num <= 0 || num > tree.Length+1 {
			return e.NewMyError("insert: line number error")
		}
		c.lineNum = num
		sliceArgs := args[2:]
		c.content = strings.Join(sliceArgs, " ")
	}
	return nil
}

func (c *insert) CallSelf() string {
	retStr := "insert"
	if c.lineNum > 0 {
		retStr += " " + strconv.Itoa(c.lineNum)
	}
	retStr += " " + c.content
	return retStr
}

type appendHead struct {
	content string
}

func (c *appendHead) Execute() (Command, error) {
	// 是否会破坏文本结构
	// 如果破坏，在哪里报错
	n, _, err := tree.ParseNode(c.content)
	if err != nil {
		return nil, err
	}
	num, err := tree.AppendHead(n)
	if err != nil {
		return nil, err
	}
	deleteC := &deleteCommand{lineNum: num}
	return deleteC, err
}

func (c *appendHead) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("append_head: args error")
	}
	sliceArgs := args[1:]
	c.content = strings.Join(sliceArgs, " ")
	return nil
}

func (c *appendHead) CallSelf() string {
	retStr := "append-head"
	retStr += " " + c.content
	return retStr
}

type appendTail struct {
	content string
}

func (c *appendTail) Execute() (Command, error) {
	n, _, err := tree.ParseNode(c.content)
	if err != nil {
		return nil, err
	}
	num, err := tree.AppendTail(n)
	if err != nil {
		return nil, err
	}
	deleteC := &deleteCommand{lineNum: num}
	return deleteC, err
}

func (c *appendTail) SetArgs(args []string) error {
	if len(args) < 2 {
		return e.NewMyError("append_tail: args error")
	}
	sliceArgs := args[1:]
	c.content = strings.Join(sliceArgs, " ")
	return nil
}

func (c *appendTail) CallSelf() string {
	retStr := "append-tail"
	retStr += " " + c.content
	return retStr
}
