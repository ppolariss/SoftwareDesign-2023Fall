package command

import (
	e "design/myError"
	"reflect"
	//"reflect"
)

// record commands which has reverse_command except undo
// record in undoHistory rather than go to canUnDoHistory when undo
// flush undoHistory when log
var canUnDoHistory []UndoableCommand

var canUnDoPointer int

type undo struct{}

func (c *undo) Execute() error {
	undoableCommand := next()
	if undoableCommand == nil {
		return nil
	}
	err := undoableCommand.UndoExecute()
	return err
}

func (c *undo) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("undo: args error")
	}
	return nil
}

func (c *undo) CallSelf() string {
	return "undo"
}

type redo struct{}

func (c *redo) Execute() error {
	command := previous()
	if command == nil {
		return nil
	}
	err := command.Execute()
	return err
}

func (c *redo) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("redo: args error")
	}
	return nil
}

func (c *redo) CallSelf() string {
	return "redo"
}

// pointer: - 1 ~ len - 1
// means next undo Command
func next() UndoableCommand {
	if canUnDoPointer >= 0 && canUnDoPointer < len(canUnDoHistory) {
		canUnDoPointer--
		return canUnDoHistory[canUnDoPointer+1]
	} else {
		return nil
	}
}

func previous() UndoableCommand {
	if canUnDoPointer >= -1 && canUnDoPointer < len(canUnDoHistory)-1 {
		canUnDoPointer++
		return canUnDoHistory[canUnDoPointer]
	} else {
		return nil
	}
}

type recordUndoableCommand struct {
}

func (c *recordUndoableCommand) update(command Command) error {
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
