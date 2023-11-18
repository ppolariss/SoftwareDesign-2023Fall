package command

import (
	e "design/myError"
	"design/util"
	"os"
	"reflect"
)

type log struct {
}

func (l *log) update(command Command) error {
	var callSelf string
	if command == nil {
		callSelf = "error"
	}
	callSelf = command.CallSelf()

	// global variable of logger
	f, err := os.OpenFile("./log/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return e.NewMyError("open log error")
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

type logFile struct {
}

func (l *logFile) update(command Command) error {
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
	f, err := os.OpenFile("./log/logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
