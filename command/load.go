package command

import (
	"errors"
)

type load struct {
	filepath string
}

func (c *load) Execute() error {
	// filepath := "../file/testFiles.txt"
	// 通过main.go运行，相对路径名要从main.go所在的目录开始！！！

	return curWorkspace.Load(c.filepath)
}

func (c *load) SetArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("load: args error")
	}
	c.filepath = args[1]
	return nil
}

func (c *load) CallSelf() string {
	return "load " + c.filepath
}
