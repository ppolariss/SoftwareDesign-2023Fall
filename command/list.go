package command

import (
	e "design/myError"
	"design/tree"
	// "fmt"
	"strings"
)

type list struct {
}

func (c *list) Execute() error {
	// tree.Dump()
	return tree.OutputAsFile(0)
}

func (c *list) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("list: args error")
	}
	return nil
}

type list_tree struct{}

func (c *list_tree) Execute() error {
	return tree.OutputAsTree()
}
func (c *list_tree) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("list_tree: args error")
	}
	return nil
}

type dir_tree struct {
	directory string
}

func (c *dir_tree) Execute() error {
	if c.directory == "" {
		return tree.OutputAsTree()
	} else {
		return tree.OutputAsDir(c.directory)
	}
}

func (c *dir_tree) SetArgs(args []string) error {
	if len(args) != 1 {
		slice_args := args[1:]
		c.directory = strings.Join(slice_args, " ")
	}
	return nil
}
