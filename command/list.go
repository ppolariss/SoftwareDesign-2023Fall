package command

import (
	"errors"
	"strings"

	"design/output"
)

type list struct {
}

func (c *list) Execute() error {
	return output.OutputAsFile(0, curWorkspace.FileContent, "")
}

func (c *list) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("list: args error")
	}
	return nil
}

func (c *list) CallSelf() string {
	return "list"
}

type listTree struct{}

func (c *listTree) Execute() error {
	return output.OutputAsTree(curWorkspace.FileContent)
}
func (c *listTree) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("list_tree: args error")
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
	if c.directory == "" {
		return output.OutputAsTree(curWorkspace.FileContent)
	} else {
		return output.OutputAsDir(c.directory, curWorkspace.FileContent)
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
