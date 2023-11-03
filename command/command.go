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
	Execute() (Command, error)
	CallSelf() string
}

type file_history struct {
	file_name string
	createAt  string
}

var commands_mapper map[string]Command
var cur_file file_history
var once sync.Once

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

// must get input outside
func Do() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		command, err := ReadCommand(scanner)
		if err != nil || command == nil {
			// if(str!="exit")
			return err
		}
		reverseCommand, err := command.Execute()
		if err != nil {
			return err
		}
		err = log(command, reverseCommand)
		if err != nil {
			return err
		}
		// logFile when save
		if reflect.TypeOf(command).Elem().Name() == "save" {
			err = logFile()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ReadCommand(scanner *bufio.Scanner) (Command, error) {
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

func log(command Command, reverseCommand Command) error {
	f, err := os.OpenFile("./log/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return e.NewMyError("open log error")
	}
	defer f.Close()

	once.Do(func() {
		util.Output("session start at "+util.GetNow()+"\n", f)
	})

	util.Output(util.GetNow()+" "+command.CallSelf()+"\n", f)
	name := reflect.TypeOf(command).Elem().Name()
	if reverseCommand != nil || name == "save" || name == "load" {
		undo_history = nil
		commands_history = append(commands_history, Command_history{command: command, reverse_command: reverseCommand})
	}
	return nil

}

func logFile() error {
	interval, err := util.GetInterval(util.GetNow(), cur_file.createAt)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("./log/logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	once.Do(func() {
		f.Truncate(0)
		util.Output("session start at "+cur_file.createAt+"\n", f)
	})
	util.Output(cur_file.file_name+" "+interval+"\n", f)
	return nil
}
