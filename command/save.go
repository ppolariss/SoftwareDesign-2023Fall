package command

import (
	e "design/myError"
	"design/output"
)

type save struct {
}

func (c *save) Execute() error {
	getFileContent()
	return output.OutputAsFile(1, fileContent, filePath)
}

func (c *save) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("save: args error")
	}
	return nil
}

func (c *save) CallSelf() string {
	return "save"
}
