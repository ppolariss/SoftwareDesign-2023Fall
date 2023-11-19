package command

import (
	"errors"

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
		return errors.New("save: args error")
	}
	return nil
}

func (c *save) CallSelf() string {
	return "save"
}
