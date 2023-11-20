package command

import (
	"design/output"
	"errors"
)

type ls struct {
}

func (c *ls) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("ls: args error")
	}
	return nil
}

func (c *ls) Execute() error {
	return output.Ls("./files/")
}
func (c *ls) CallSelf() string {
	return "ls"
}
