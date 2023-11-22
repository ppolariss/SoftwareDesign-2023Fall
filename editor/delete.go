package editor

import (
	"design/util"
	"design/workspace"
	"errors"
	"strconv"
	"strings"
)

type DeleteCommand struct {
	lineNum int
	content string //node.content
}

func (c *DeleteCommand) Execute() error {
	var err error
	c.lineNum, c.content, err = Delete(c.lineNum, c.content, &workspace.CurWorkspace.FileContent)
	if err != nil {
		return err
	}
	return nil
	// 删除指定标题或⽂本。如果指定⾏号，则删除指定⾏。当删除的是标题时，其⼦标题
	// 和内容不会被删除。
}

func (c *DeleteCommand) SetArgs(args []string) error {
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

func (c *DeleteCommand) CallSelf() string {
	retStr := "delete"
	if c.lineNum > 0 {
		retStr += " " + strconv.Itoa(c.lineNum)
	} else {
		retStr += " " + c.content
	}
	return retStr
}

func (c *DeleteCommand) UndoExecute() error {
	i := InsertCommand{lineNum: c.lineNum, content: c.content}
	return i.Execute()
}

// Delete para line_num or content
// return index content error
func Delete(lineNum int, content string, fileContent *[]string) (int, string, error) {

	var originContent string
	if !util.IsInit() {
		return 0, "", errors.New("delete: No file in workspace")
	}
	if lineNum > 0 {
		if lineNum > len(*fileContent) {
			return 0, "", errors.New("delete: line number out of range")
		}
		originContent = (*fileContent)[lineNum-1]
	} else {
		var err error
		lineNum, err = matchContent(content, *fileContent)
		if err != nil {
			return 0, "", err
		}
		if lineNum <= 0 || lineNum > len(*fileContent) {
			return 0, "", errors.New("delete: content not match")
		}
		originContent = (*fileContent)[lineNum-1]
	}
	*fileContent = append((*fileContent)[:lineNum-1], (*fileContent)[lineNum:]...)

	return lineNum, originContent, nil
}
