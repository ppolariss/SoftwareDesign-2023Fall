package commandManager

import (
	"errors"
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
		return errors.New("undo: args error")
	}
	return nil
}

func (c *Undo) CallSelf() string {
	return "undo"
}
