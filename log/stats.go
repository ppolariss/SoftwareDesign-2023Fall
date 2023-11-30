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
	if workspace.CurWorkspace == nil && c.status == "current" {
		return errors.New("stats: curWorkspace is nil")
	}
	return stats(c.status, workspace.CurWorkspace.CreateAt, workspace.CurWorkspace.FileName)
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

func stats(status string, createAt string, fileName string) error {
	if status == "current" {
		interval, err := util.GetInterval(util.GetNow(), createAt)
		if err != nil {
			return err
		}
		fmt.Println(fileName + " " + interval)
		return nil
	}
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE, 0644)
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
