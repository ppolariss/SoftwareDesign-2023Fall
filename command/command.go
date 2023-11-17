package command

import (
	"bufio"
	e "design/myError"
	"os"
	"reflect"

	"design/util"
	"strings"
	"sync"
)

type Command interface {
	SetArgs([]string) error
	Execute() error
	CallSelf() string
}

type UndoableCommand interface {
	Command
	UndoExecute() error
}

type fileHistory struct {
	fileName string
	createAt string
}

var commandsMapper map[string]Command
var curFile fileHistory
var once sync.Once

func init() {
	commandsMapper = make(map[string]Command)
	commandsMapper["load"] = &load{}
	commandsMapper["save"] = &save{}
	commandsMapper["insert"] = &insert{}
	commandsMapper["delete"] = &deleteCommand{}
	commandsMapper["append-head"] = &appendHead{}
	commandsMapper["append-tail"] = &appendTail{}
	commandsMapper["undo"] = &undo{}
	commandsMapper["redo"] = &redo{}
	commandsMapper["list"] = &list{}
	commandsMapper["list-tree"] = &listTree{}
	commandsMapper["dir-tree"] = &dirTree{}
	commandsMapper["history"] = &history{}
	commandsMapper["stats"] = &stats{}
}

// Do must get input outside
func Do() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		command, err := ReadCommand(scanner)
		if err != nil || command == nil {
			// if(str!="exit")
			return err
		}
		err = command.Execute()
		if err != nil {
			return err
		}
		err = log(command)
		if err != nil {
			return err
		}
		updateCanUndoHistory(command)

		// logFile when save
		if reflect.TypeOf(command).Elem().Name() == "save" {
			err = logFile()
			if err != nil {
				return err
			}
		}
	}
	// 错误日志
	return nil
}

func ReadCommand(scanner *bufio.Scanner) (Command, error) {
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, e.NewMyError("Input error")
	}
	// parse input to args
	splitStrings := strings.Split(line, " ")
	args := make([]string, 0)
	for _, str := range splitStrings {
		if str != "" {
			args = append(args, str)
		}
	}

	// fmt.Println(args)
	if len(args) == 0 {
		return nil, e.NewMyError("Null input")
	}
	// get command
	command := commandsMapper[args[0]]
	if command == nil {
		return nil, e.NewMyError("invalid command")
	}
	err := command.SetArgs(args)
	if err != nil {
		return nil, err
	}
	return command, nil
}

func log(command Command) error {
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

	_ = util.Output(util.GetNow()+" "+command.CallSelf()+"\n", f)
	return nil
}

func updateCanUndoHistory(command Command) {
	name := reflect.TypeOf(command).Elem().Name()
	if name == "save" || name == "load" {
		canUnDoHistory = canUnDoHistory[:0]
		canUnDoPointer = 0
		return
	}
	if undoableCommand, ok := command.(UndoableCommand); ok {
		canUnDoHistory = append(canUnDoHistory, undoableCommand)
		canUnDoPointer = len(canUnDoHistory) - 1
	}
}

func logFile() error {
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
