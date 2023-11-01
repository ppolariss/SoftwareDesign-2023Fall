package command

import (
	"fmt"
	e "design/myError"
	"strconv"
)


type undo struct {
}

func (c *undo) Execute() error {
	fmt.Println("undo")
	return nil
}
func (c *undo) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("undo: args error")
	}
	return nil
}

type redo struct{}

func (c *redo) Execute() error {
	fmt.Println("redo")
	return nil
}
func (c *redo) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("redo: args error")
	}
	return nil
}


type history struct {
	// default -1
	num int
}

func (c *history) Execute() error {
	fmt.Println("history")
	return nil
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

type stats struct {
	// default current
	status string
}

func (c *stats) Execute() error {
	fmt.Println("stats")
	return nil
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