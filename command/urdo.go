package command

import (
	e "design/myError"
	"reflect"
)

type History struct {
	command        Command
	reverseCommand Command
}

// record commands which has reverse_command except undo
var commandsHistory []History

// record in undoHistory rather than go to commandsHistory when undo
// flush undoHistory when log
var undoHistory *History

type undo struct{}

func (c *undo) Execute() (Command, error) {
	ch := next()
	if ch == nil {
		return nil, nil
	}
	command, err := ch.reverseCommand.Execute()
	if err != nil {
		return nil, err
	}
	undoHistory = &History{command: c, reverseCommand: command}
	return nil, nil
	// 因为只有redo会检测undo所以希望undo不进入next
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

func (c *redo) Execute() (Command, error) {
	if undoHistory == nil {
		return nil, nil
	}
	if reflect.TypeOf(undoHistory.command).Elem().Name() != "undo" {
		return nil, nil
	}
	command, err := undoHistory.reverseCommand.Execute()
	if err != nil {
		return nil, err
	}
	undoHistory = nil
	return command, nil
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

func next() *History {
	pointer := len(commandsHistory) - 1
	// delete after undo
	// insert after redo
	// simulate stack
	for {
		if pointer < 0 {
			// !!!!!!!!! condition isn't equal 0
			return nil
		}
		c := commandsHistory[pointer]

		if c.reverseCommand == nil {
			name := reflect.TypeOf(c.command).Elem().Name()
			if name == "save" || name == "load" {
				return nil
			}
			pointer--
			commandsHistory = commandsHistory[:pointer+1]
			continue
		}

		commandsHistory = commandsHistory[:pointer]
		return &c
	}
}
