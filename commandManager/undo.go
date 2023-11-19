package commandManager

import (
	e "design/myError"
	//"reflect"
)

type Undo struct{}

func (c *Undo) Execute() error {
	undoableCommand := next()
	if undoableCommand == nil {
		return nil
	}
	err := undoableCommand.UndoExecute()
	return err
}

func (c *Undo) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("undo: args error")
	}
	return nil
}

func (c *Undo) CallSelf() string {
	return "undo"
}
