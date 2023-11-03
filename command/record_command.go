package command

import (
	e "design/myError"
	"fmt"
	"os"

	"design/util"
	"strconv"
)


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
