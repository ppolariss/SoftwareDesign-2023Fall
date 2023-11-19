package command

import (
	"errors"
	"strings"

	"design/editor"
	"design/util"
)

var filePath string

type load struct {
	filepath string
}

func (c *load) Execute() error {
	// filepath := "../file/testFiles.txt"
	// 通过main.go运行，相对路径名要从main.go所在的目录开始！！！
	if strings.HasSuffix(c.filepath, "file/") {
		c.filepath = "./" + c.filepath
	} else {
		c.filepath = "./file/" + c.filepath
	}

	curFile.fileName = c.filepath
	filePath = c.filepath
	curFile.createAt = util.GetNow()
	return editor.Load(c.filepath)
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
