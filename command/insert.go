package command

import (
	"design/editor"
	"design/workspace"
	"errors"
	"strconv"
	"strings"
)

type InsertCommand struct {
	Name string
	// -1 表示最后
	LineNum int
	Content string
}

func (c *InsertCommand) Execute() error {
	var err error
	c.LineNum, err = editor.Insert(c.LineNum, c.Content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *InsertCommand) SetArgs(args []string) error {
	c.Name = "insert"
	if len(args) < 2 {
		return errors.New("insert: args error")
	} else if len(args) == 2 {
		c.LineNum = -1
		c.Content = args[1]
	} else {
		num, err := strconv.Atoi(args[1])
		// The line number is not specified
		if err != nil {
			c.LineNum = -1
			sliceArgs := args[1:]
			c.Content = strings.Join(sliceArgs, " ")
			return nil
		}
		// The line number is specified
		// in order not to couple the bottom
		//
		if num <= 0 || num > len(workspace.CurWorkspace.FileContent)+1 {
			return errors.New("insert: line number error")
		}
		c.LineNum = num
		sliceArgs := args[2:]
		c.Content = strings.Join(sliceArgs, " ")
	}
	return nil
}

func (c *InsertCommand) CallSelf() string {
	retStr := "insert"
	if c.LineNum > 0 {
		retStr += " " + strconv.Itoa(c.LineNum)
	}
	retStr += " " + c.Content
	return retStr
}

func (c *InsertCommand) UndoExecute() error {
	command := DeleteCommand{LineNum: c.LineNum, Content: c.Content}
	return command.Execute()
}

type AppendHead struct {
	Name    string
	Content string
}

func (c *AppendHead) Execute() error {
	// 是否会破坏文本结构
	// 如果破坏，在哪里报错

	_, err := editor.Insert(1, c.Content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *AppendHead) SetArgs(args []string) error {
	c.Name = "append-head"
	if len(args) < 2 {
		return errors.New("append_head: args error")
	}
	sliceArgs := args[1:]
	c.Content = strings.Join(sliceArgs, " ")
	return nil
}

func (c *AppendHead) CallSelf() string {
	return "append-head " + c.Content
}

func (c *AppendHead) UndoExecute() error {
	command := DeleteCommand{LineNum: 1, Content: c.Content}
	return command.Execute()
}

type AppendTail struct {
	Name    string
	LineNum int
	Content string
}

func (c *AppendTail) Execute() error {
	c.Name = "append-tail"
	var err error
	c.LineNum, err = editor.Insert(-1, c.Content, &workspace.CurWorkspace.FileContent)
	return err
}

func (c *AppendTail) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("append_tail: args error")
	}
	sliceArgs := args[1:]
	c.Content = strings.Join(sliceArgs, " ")
	return nil
}

func (c *AppendTail) CallSelf() string {
	return "append-tail " + c.Content
}

func (c *AppendTail) UndoExecute() error {
	command := DeleteCommand{LineNum: c.LineNum, Content: c.Content}
	return command.Execute()
}
