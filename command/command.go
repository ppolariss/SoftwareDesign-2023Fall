package command

import (
	"bufio"
	e "design/myError"
	"os"
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

	registerObserver(&recordUndoableCommand{})
	registerObserver(&log{})
	registerObserver(&logFile{})
}

// Do must get input outside
func Do() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var err error
	for scanner.Scan() {
		command, err := ReadCommand(scanner)
		if err != nil || command == nil {
			// if(str!="exit")
			break
		}
		err = command.Execute()
		if err != nil {
			break
		}

		err = notifyObserver(command)
		if err != nil {
			return err
		}
	}
	if err != nil {
		// 错误日志
		err = notifyObserver(nil)
	}
	return err
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
