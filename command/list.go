package command

import (
	e "design/myError"
	"design/tree"
	"strings"
)

type list struct {
}

func (c *list) Execute() (Command, error) {
	// tree.Dump()
	return nil, tree.OutputAsFile(0)
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

func (c *listTree) Execute() (Command, error) {
	return nil, tree.OutputAsTree()
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

func (c *dirTree) Execute() (Command, error) {
	if c.directory == "" {
		return nil, tree.OutputAsTree()
	} else {
		return nil, tree.OutputAsDir(c.directory)
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
