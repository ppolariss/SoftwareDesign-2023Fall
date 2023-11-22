package log

import (
	"design/util"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type History struct {
	// default -1
	num int
}

func (c *History) Execute() error {
	return history(c.num)
}

func (c *History) SetArgs(args []string) error {
	if len(args) > 2 {
		return errors.New("history: args error")
	}
	if len(args) == 1 {
		c.num = -1
	} else {
		num, err := strconv.Atoi(args[1])
		if err != nil || num < 1 {
			return errors.New("history: args error")
		}
		c.num = num
	}
	return nil
}

func (c *History) CallSelf() string {
	retStr := "history"
	if c.num > 0 {
		retStr += " " + strconv.Itoa(c.num)
	}
	return retStr
}

func history(num int) error {
	f, err := os.OpenFile("./logFiles/log", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	tempHistory, err := util.ReadStrings(f)
	if err != nil {
		return err
	}

	var count int
	if num == -1 {
		// count = len(commands_history)
		count = len(tempHistory)
	} else {
		count = num
	}

	for i := len(tempHistory); i > 0 && count > 0; i-- {
		fmt.Println(tempHistory[i-1])
		count--
	}
	return nil
}
