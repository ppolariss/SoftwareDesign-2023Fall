package command

import (
	. "design/interfaces"
	e "design/myError"
	"design/util"
	"fmt"
	"os"
	"reflect"
	"sync"
)

var once sync.Once

type stats struct {
	// default current
	status string
}

func (c *stats) Execute() error {
	if c.status == "current" {
		interval, err := util.GetInterval(util.GetNow(), curFile.createAt)
		if err != nil {
			return err
		}
		fmt.Println(curFile.fileName + " " + interval)
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

type logFile struct {
}

func (l *logFile) Update(command Command) error {
	if command == nil {
		return nil
	}
	if reflect.TypeOf(command).Elem().Name() != "save" {
		return nil
	}

	interval, err := util.GetInterval(util.GetNow(), curFile.createAt)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("./logFiles/logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	once.Do(func() {
		_ = f.Truncate(0)
		_ = util.Output("session start at "+curFile.createAt+"\n", f)
	})
	_ = util.Output(curFile.fileName+" "+interval+"\n", f)
	return nil
}
