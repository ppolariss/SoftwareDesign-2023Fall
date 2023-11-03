package command

import (
	e "design/myError"
	"design/tree"
)

type save struct {
}

func (c *save) Execute() (Command, error) {
	return nil, tree.OutputAsFile(1)
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
