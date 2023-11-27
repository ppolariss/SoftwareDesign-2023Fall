package log

import (
	"design/workspace"
	"errors"
	"os"
	"sync"

	. "design/interfaces"
	"design/util"
)

var once sync.Once

type Log struct {
}

func (l *Log) Update(command Command) error {
	var fileName string
	if workspace.CurWorkspace == nil || workspace.CurWorkspace.FileName == "" {
		fileName = "global"
	} else {
		fileName = workspace.CurWorkspace.FileName
	}

	var callSelf string
	if command == nil {
		callSelf = "error"
	} else {
		callSelf = command.CallSelf()
	}

	// global variable of logger
	// 可能要转义
	f, err := os.OpenFile("./logFiles/log/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("open log error")
	}
	defer func() {
		_ = f.Close()
	}()

	once.Do(func() {
		_ = util.Output("session start at "+util.GetNow()+"\n", f)
	})

	_ = util.Output(util.GetNow()+" "+callSelf+"\n", f)
	return nil
}
