package log

import (
	"design/util"
	"design/workspace"
	"errors"
	"fmt"
	"os"
)

type Stats struct {
	// default current
	status string
}

func (c *Stats) Execute() error {
	if c.status == "current" {
		interval, err := util.GetInterval(util.GetNow(), workspace.CurWorkspace.CreateAt)
		if err != nil {
			return err
		}
		fmt.Println(workspace.CurWorkspace.FileName + " " + interval)
		return nil
	}
	f, err := os.OpenFile("./logFiles/logFile", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	tempHistory, err := util.ReadStrings(f)
	if err != nil {
		return nil
	}
	for _, v := range tempHistory {
		fmt.Println(v)
	}
	return nil
}

func (c *Stats) SetArgs(args []string) error {
	if len(args) > 2 {
		return errors.New("stats: args error")
	}
	if len(args) == 1 {
		c.status = "current"
	} else {
		if args[1] != "current" && args[1] != "all" {
			return errors.New("stats: args error")
		}
		c.status = args[1]
	}
	return nil
}

func (c *Stats) CallSelf() string {
	return "stats" + " " + c.status
}
