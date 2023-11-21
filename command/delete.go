package command

import (
	"errors"

	"design/editor"
	// "fmt"
	"strconv"
	"strings"
)

type deleteCommand struct {
	lineNum int
	content string //node.content
}

func (c *deleteCommand) Execute() error {
	var err error
	c.lineNum, c.content, err = editor.Delete(c.lineNum, c.content, &curWorkspace.FileContent)
	if err != nil {
		return err
	}
	return nil
	// 删除指定标题或⽂本。如果指定⾏号，则删除指定⾏。当删除的是标题时，其⼦标题
	// 和内容不会被删除。
}

func (c *deleteCommand) SetArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("delete: args error")
	}

	// parse to line_num if and only if the rest is a positive number
	if len(args) == 2 {
		num, err := strconv.Atoi(args[1])
		if err == nil && num > 0 {
			if num > len(curWorkspace.FileContent) {
				return errors.New("delete: line number error")
			}
			c.lineNum = num
			c.content = ""
			return nil
		}
	}

	// The line number is not specified
	c.lineNum = 0
	sliceArgs := args[1:]
	c.content = strings.Join(sliceArgs, " ")

	return nil
}

func (c *deleteCommand) CallSelf() string {
	retStr := "delete"
	if c.lineNum > 0 {
		retStr += " " + strconv.Itoa(c.lineNum)
	} else {
		retStr += " " + c.content
	}
	return retStr
}

func (c *deleteCommand) UndoExecute() error {
	i := insert{lineNum: c.lineNum, content: c.content}
	return i.Execute()
}
