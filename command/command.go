package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"design/commandManager"
	. "design/interfaces"
	"design/log"
)

func init() {
	RegisterObserver(&commandManager.RecordUndoableCommand{})
	RegisterObserver(&log.Log{})
	RegisterObserver(&log.LogFile{})
}

// Do must get input outside
func Do() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	var err error
	for scanner.Scan() {
		var command Command
		command, err = ReadCommand(scanner)
		if err != nil || command == nil {
			// if(str!="exit")
			fmt.Println("invalid command")
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
		_ = NotifyObserver(nil)
	}
	return err
}

func ReadCommand(scanner *bufio.Scanner) (Command, error) {
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, errors.New("input error")
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
		return nil, errors.New("no input")
	}
	// get command
	command := commandsMapper[args[0]]
	if command == nil {
		return nil, errors.New("invalid command")
	}
	err := command.SetArgs(args)
	if err != nil {
		return nil, err
	}
	return command, nil
}
