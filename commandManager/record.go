package commandManager

import (
	. "design/interfaces"
	"reflect"
)

type RecordUndoableCommand struct {
}

func (c *RecordUndoableCommand) Update(command Command) error {
	if command == nil {
		return nil
	}
	name := reflect.TypeOf(command).Elem().Name()
	if name == "save" || name == "load" {
		canUnDoHistory = canUnDoHistory[:0]
		canUnDoPointer = 0
		return nil
	}
	// if command is undoable
	undoableCommand, ok := command.(UndoableCommand)
	if ok {
		canUnDoHistory = append(canUnDoHistory, undoableCommand)
		canUnDoPointer = len(canUnDoHistory) - 1
	}
	return nil
}
