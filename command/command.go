package command

import (
	"bufio"
	"design/commandManager"
	. "design/interfaces"
	"design/log"
	e "design/myError"
	"os"
	"strings"
)

type fileHistory struct {
	fileName string
	createAt string
}

var commandsMapper map[string]Command
var curFile fileHistory

func init() {
	commandsMapper = make(map[string]Command)
	commandsMapper["load"] = &load{}
	commandsMapper["save"] = &save{}
	commandsMapper["insert"] = &insert{}
	commandsMapper["delete"] = &deleteCommand{}
	commandsMapper["append-head"] = &appendHead{}
	commandsMapper["append-tail"] = &appendTail{}
	commandsMapper["undo"] = &commandManager.Undo{}
	commandsMapper["redo"] = &commandManager.Redo{}
	commandsMapper["list"] = &list{}
	commandsMapper["list-tree"] = &listTree{}
	commandsMapper["dir-tree"] = &dirTree{}
	commandsMapper["history"] = &log.History{}
	commandsMapper["stats"] = &stats{}
	commandsMapper["ls"] = &ls{}

	RegisterObserver(&commandManager.RecordUndoableCommand{})
	RegisterObserver(&log.Log{})
	RegisterObserver(&logFile{})
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

		err = NotifyObserver(command)
		if err != nil {
			return err
		}
	}
	if err != nil {
		// 错误日志
		err = NotifyObserver(nil)
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
