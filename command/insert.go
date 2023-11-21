package command

import (
	"design/workspace"
	"errors"

	"design/editor"
	// "fmt"
	"strconv"
	"strings"
)

type insert struct {
	// -1 表示最后
	lineNum int
	content string
}

func (c *insert) Execute() error {
	var err error
	c.lineNum, err = editor.Insert(c.lineNum, c.content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *insert) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("insert: args error")
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
		// in order not to couple the bottom
		//
		if num <= 0 || num > len(workspace.CurWorkspace.FileContent)+1 {
			return errors.New("insert: line number error")
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

func (c *insert) UndoExecute() error {
	command := deleteCommand{lineNum: c.lineNum, content: c.content}
	return command.Execute()
}

type appendHead struct {
	content string
}

func (c *appendHead) Execute() error {
	// 是否会破坏文本结构
	// 如果破坏，在哪里报错

	_, err := editor.Insert(1, c.content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *appendHead) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("append_head: args error")
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

func (c *appendHead) UndoExecute() error {
	command := deleteCommand{lineNum: 1, content: c.content}
	return command.Execute()
}

type appendTail struct {
	lineNum int
	content string
}

func (c *appendTail) Execute() error {
	var err error
	c.lineNum, err = editor.Insert(-1, c.content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *appendTail) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("append_tail: args error")
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

func (c *appendTail) UndoExecute() error {
	command := deleteCommand{lineNum: c.lineNum, content: c.content}
	return command.Execute()
}
