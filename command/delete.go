package command

import (
	"design/editor"
	"design/workspace"
	"errors"
	"strconv"
	"strings"
)

type DeleteCommand struct {
	Name    string
	LineNum int
	Content string //node.Content
}

func (c *DeleteCommand) Execute() error {
	if workspace.CurWorkspace == nil {
		return errors.New("delete: no workspace")
	}
	var err error
	c.LineNum, c.Content, err = editor.Delete(c.LineNum, c.Content, &workspace.CurWorkspace.FileContent)
	return err
	// 删除指定标题或⽂本。如果指定⾏号，则删除指定⾏。当删除的是标题时，其⼦标题
	// 和内容不会被删除。
}

func (c *DeleteCommand) SetArgs(args []string) error {
	c.Name = "delete"
	if len(args) < 2 {
		return errors.New("delete: args error")
	}

	// parse to line_num if and only if the rest is a positive number
	if len(args) == 2 {
		num, err := strconv.Atoi(args[1])
		if err == nil && num > 0 {
			if num > len(workspace.CurWorkspace.FileContent) {
				return errors.New("delete: line number error")
			}
			c.LineNum = num
			c.Content = ""
			return nil
		}
	}

	// The line number is not specified
	c.LineNum = 0
	sliceArgs := args[1:]
	c.Content = strings.Join(sliceArgs, " ")

	return nil
}

func (c *DeleteCommand) CallSelf() string {
	retStr := "delete"
	if c.LineNum > 0 {
		retStr += " " + strconv.Itoa(c.LineNum)
	} else {
		retStr += " " + c.Content
	}
	return retStr
}

func (c *DeleteCommand) UndoExecute() error {
	i := InsertCommand{LineNum: c.LineNum, Content: c.Content}
	return i.Execute()
}
