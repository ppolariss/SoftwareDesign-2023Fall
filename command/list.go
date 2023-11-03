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

type list_tree struct{}

func (c *list_tree) Execute() (Command, error) {
	return nil, tree.OutputAsTree()
}
func (c *list_tree) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("list_tree: args error")
	}
	return nil
}
func (c *list_tree) CallSelf() string {
	return "list-tree"
}

type dir_tree struct {
	directory string
}

func (c *dir_tree) Execute() (Command, error) {
	if c.directory == "" {
		return nil, tree.OutputAsTree()
	} else {
		return nil, tree.OutputAsDir(c.directory)
	}
}

func (c *dir_tree) SetArgs(args []string) error {
	if len(args) != 1 {
		slice_args := args[1:]
		c.directory = strings.Join(slice_args, " ")
	}
	return nil
}
func (c *dir_tree) CallSelf() string {
	return "dir-tree"
}
