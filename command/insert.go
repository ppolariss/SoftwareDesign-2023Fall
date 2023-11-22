package command

import (
	"design/editor"
	"design/workspace"
	"errors"
	"strconv"
	"strings"
)

type InsertCommand struct {
	// -1 表示最后
	lineNum int
	content string
}

func (c *InsertCommand) Execute() error {
	var err error
	c.lineNum, err = editor.Insert(c.lineNum, c.content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *InsertCommand) SetArgs(args []string) error {
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

func (c *InsertCommand) CallSelf() string {
	retStr := "insert"
	if c.lineNum > 0 {
		retStr += " " + strconv.Itoa(c.lineNum)
	}
	retStr += " " + c.content
	return retStr
}

func (c *InsertCommand) UndoExecute() error {
	command := DeleteCommand{lineNum: c.lineNum, content: c.content}
	return command.Execute()
}

type AppendHead struct {
	content string
}

func (c *AppendHead) Execute() error {
	// 是否会破坏文本结构
	// 如果破坏，在哪里报错

	_, err := editor.Insert(1, c.content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *AppendHead) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("append_head: args error")
	}
	sliceArgs := args[1:]
	c.content = strings.Join(sliceArgs, " ")
	return nil
}

func (c *AppendHead) CallSelf() string {
	retStr := "append-head"
	retStr += " " + c.content
	return retStr
}

func (c *AppendHead) UndoExecute() error {
	command := DeleteCommand{lineNum: 1, content: c.content}
	return command.Execute()
}

type AppendTail struct {
	lineNum int
	content string
}

func (c *AppendTail) Execute() error {
	var err error
	c.lineNum, err = editor.Insert(-1, c.content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *AppendTail) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("append_tail: args error")
	}
	sliceArgs := args[1:]
	c.content = strings.Join(sliceArgs, " ")
	return nil
}

func (c *AppendTail) CallSelf() string {
	retStr := "append-tail"
	retStr += " " + c.content
	return retStr
}

func (c *AppendTail) UndoExecute() error {
	command := DeleteCommand{lineNum: c.lineNum, content: c.content}
	return command.Execute()
}
