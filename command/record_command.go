package command

import (
	e "design/myError"
	"fmt"
	"os"

	"design/util"
	"reflect"
	"strconv"
)

var undo_history *Command_history

type undo struct{}

func (c *undo) Execute() (Command, error) {
	ch := next()
	if ch == nil {
		return nil, nil
	}
	command, err := ch.reverse_command.Execute()
	if err != nil {
		return nil, err
	}
	undo_history = &Command_history{command: c, reverse_command: command}
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
	if undo_history == nil {
		return nil, nil
	}
	if reflect.TypeOf(undo_history.command).Elem().Name() != "undo" {
		return nil, nil
	}
	command, err := undo_history.reverse_command.Execute()
	if err != nil {
		return nil, err
	}
	undo_history = nil
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

type history struct {
	// default -1
	num int
}

func (c *history) Execute() (Command, error) {
	f, err := os.OpenFile("./log/log", os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	temp_history, err := util.ReadStrings(f)
	if err != nil {
		return nil, err
	}

	var count int
	if c.num == -1 {
		// count = len(commands_history)
		count = len(temp_history)
	} else {
		count = c.num
	}

	for i := len(temp_history); i > 0 && count > 0; i-- {
		fmt.Println(temp_history[i-1])
		count--
	}

	return nil, nil
}

func (c *history) SetArgs(args []string) error {
	if len(args) > 2 {
		return e.NewMyError("history: args error")
	}
	if len(args) == 1 {
		c.num = -1
	} else {
		num, err := strconv.Atoi(args[1])
		if err != nil || num < 1 {
			return e.NewMyError("history: args error")
		}
		c.num = num
	}
	return nil
}

func (c *history) CallSelf() string {
	retStr := "history"
	if c.num > 0 {
		retStr += " " + strconv.Itoa(c.num)
	}
	return retStr
}

type stats struct {
	// default current
	status string
}

func (c *stats) Execute() (Command, error) {
	if c.status == "current" {
		interval, err := util.GetInterval(util.GetNow(), cur_file.createAt)
		if err != nil {
			return nil, err
		}
		fmt.Println(cur_file.file_name + " " + interval)
		return nil, nil
	}
	f, err := os.OpenFile("./log/logFile", os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	temp_history, err := util.ReadStrings(f)
	if err != nil {
		return nil, err
	}
	for _, v := range temp_history {
		fmt.Println(v)
	}
	return nil, nil
}

func (c *stats) SetArgs(args []string) error {
	if len(args) > 2 {
		return e.NewMyError("stats: args error")
	}
	if len(args) == 1 {
		c.status = "current"
	} else {
		if args[1] != "current" && args[1] != "all" {
			return e.NewMyError("stats: args error")
		}
		c.status = args[1]
	}
	return nil
}

func (c *stats) CallSelf() string {
	return "stats" + " " + c.status
}

func next() *Command_history {
	pointer := len(commands_history) - 1
	// delete after undo
	// insert after redo
	// simulate stack
	for {
		if pointer < 0 {
			// !!!!!!!!! condition isn't equal 0
			return nil
		}
		c := commands_history[pointer]

		if c.reverse_command == nil {
			name := reflect.TypeOf(c.command).Elem().Name()
			if name == "save" || name == "load" {
				return nil
			}
			pointer--
			commands_history = commands_history[:pointer+1]
			continue
		}

		commands_history = commands_history[:pointer]
		return &c
	}
}
