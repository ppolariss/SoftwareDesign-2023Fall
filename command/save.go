package command

import (
	"design/workspace"
	"errors"

	"design/output"
)

type save struct {
}

func (c *save) Execute() error {
	return output.OutputAsFile(1, curWorkspace.FileContent, workspace.GetFilePath(curWorkspace.FileName))
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
