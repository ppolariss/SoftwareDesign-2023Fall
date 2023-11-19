package commandManager

import (
	e "design/myError"
)

type Redo struct{}

func (c *Redo) Execute() error {
	command := previous()
	if command == nil {
		return nil
	}
	err := command.Execute()
	return err
}

func (c *Redo) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("redo: args error")
	}
	return nil
}

func (c *Redo) CallSelf() string {
	return "redo"
}
