package command

import (
	"design/output"
	"design/workspace"
	"errors"
	"strings"
)

type List struct {
}

func (c *List) Execute() error {
	return output.AsFile(0, workspace.CurWorkspace.FileContent, "")
}

func (c *List) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("list: args error")
	}
	return nil
}

func (c *List) CallSelf() string {
	return "list"
}

type ListTree struct{}

func (c *ListTree) Execute() error {
	return output.AsTree(workspace.CurWorkspace.FileContent)
}
func (c *ListTree) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("list-tree: args error")
	}
	return nil
}
func (c *ListTree) CallSelf() string {
	return "list-tree"
}

type DirTree struct {
	directory string
}

func (c *DirTree) Execute() error {
	if c.directory == "" {
		return output.AsTree(workspace.CurWorkspace.FileContent)
	} else {
		return output.AsDir(c.directory, workspace.CurWorkspace.FileContent)
	}
}

func (c *DirTree) SetArgs(args []string) error {
	if len(args) != 1 {
		sliceArgs := args[1:]
		c.directory = strings.Join(sliceArgs, " ")
	} else {
		return errors.New("dir-tree: args error")
	}
	return nil
}
func (c *DirTree) CallSelf() string {
	return "dir-tree"
}
