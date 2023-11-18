package command

import (
	"design/fileEditor"
	e "design/myError"
	"design/output"
	"strings"
)

type list struct {
}

var fileContent []string

func (c *list) Execute() error {
	// tree.Dump()
	getFileContent()
	return output.OutputAsFile(0, fileContent, filePath)
}

func (c *list) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("list: args error")
	}
	return nil
}

func (c *list) CallSelf() string {
	return "list"
}

type listTree struct{}

func (c *listTree) Execute() error {
	getFileContent()
	return output.OutputAsTree(fileContent)
}
func (c *listTree) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("list_tree: args error")
	}
	return nil
}
func (c *listTree) CallSelf() string {
	return "list-tree"
}

type dirTree struct {
	directory string
}

func (c *dirTree) Execute() error {
	getFileContent()
	if c.directory == "" {
		return output.OutputAsTree(fileContent)
	} else {
		return output.OutputAsDir(c.directory, fileContent)
	}
}

func (c *dirTree) SetArgs(args []string) error {
	if len(args) != 1 {
		sliceArgs := args[1:]
		c.directory = strings.Join(sliceArgs, " ")
	}
	return nil
}
func (c *dirTree) CallSelf() string {
	return "dir-tree"
}

func getFileContent() {
	fileContent = fileEditor.GetFileContent()
}
