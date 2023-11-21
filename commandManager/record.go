package commandManager

import (
	. "design/interfaces"
	. "design/workspace"
)

type RecordUndoableCommand struct {
}

func (c *RecordUndoableCommand) Update(command Command) error {
	if command == nil {
		return nil
	}

	// if command is undoable
	undoableCommand, ok := command.(UndoableCommand)
	if ok {
		CurWorkspace.UndoableCommandHistory = append(CurWorkspace.UndoableCommandHistory, undoableCommand)
		CurWorkspace.UndoableCommandPointer = len(CurWorkspace.UndoableCommandHistory) - 1
	}
	return nil
}
