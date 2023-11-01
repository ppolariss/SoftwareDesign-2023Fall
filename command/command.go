package command

import (
	// "context"
	"bufio"
	e "design/myError"
	"fmt"
	"os"
	// "strconv"
	"strings"
	// "golang.org/x/exp/slices"
)

type Command interface {
	SetArgs([]string) error
	Execute() error
}

var commands_mapper map[string]Command

func init() {
	commands_mapper = make(map[string]Command)
	commands_mapper["load"] = &load{}
	commands_mapper["save"] = &save{}
	commands_mapper["insert"] = &insert{}
	commands_mapper["delete"] = &delete{}
	commands_mapper["append-head"] = &append_head{}
	commands_mapper["append-tail"] = &append_tail{}
	commands_mapper["undo"] = &undo{}
	commands_mapper["redo"] = &redo{}
	commands_mapper["list"] = &list{}
	commands_mapper["list-tree"] = &list_tree{}
	commands_mapper["dir-tree"] = &dir_tree{}
	commands_mapper["history"] = &history{}
	commands_mapper["stats"] = &stats{}
}

func Do() error {
	// for{}
	command, err := ReadCommand()
	if err != nil || command == nil {
		// if(str!="exit")
		fmt.Println(err)
		// panic("")
		return err
	}
	err = command.Execute()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ReadCommand() (Command, error) {
	// get input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, e.NewMyError("Input error")
	}
	// parse input to args
	split_strings := strings.Split(line, " ")
	args := make([]string, 0)
	for _, str := range split_strings {
		if str != "" {
			args = append(args, str)
		}
	}

	// fmt.Println(args)
	if len(args) == 0 {
		return nil, e.NewMyError("Null input")
	}
	// get command
	command := commands_mapper[args[0]]
	if command == nil {
		return nil, e.NewMyError("invalid command")
	}
	err := command.SetArgs(args)
	if err != nil {
		return nil, err
	}
	return command, nil
}
