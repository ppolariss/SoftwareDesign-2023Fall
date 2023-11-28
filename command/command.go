package command

import (
	"design/util"
	"design/workspace"
	"errors"
	"strings"

	"design/commandManager"
	. "design/interfaces"
	"design/log"
)

func Init() {
	RegisterObserver(&commandManager.RecordUndoableCommand{})
	RegisterObserver(&log.Log{})
	Deserialize()
	_ = util.AsJson("", workspace.Path+"backup.json")
	//RegisterObserver(&log.LogFile{})
}

// Do must get input outside
func Do() error {
	//reader := bufio.NewReader(os.Stdin)
	//reader.ReadLine()
	var err error
	//input := util.GetInput()
	//line := range input
	for {
		var line string
		line, err = util.GetInput()
		if err != nil {
			return err
		}
		if line == "" {
			return errors.New("no command input")
		}
		var command Command
		command, err = ReadCommand(line)
		if err != nil || command == nil {
			// if(str!="exit")
			//fmt.Println("invalid command")
			return errors.New("invalid command")
		}
		err = command.Execute()

		if err != nil {
			if err.Error() == "exit" {
				return NotifyObserver(command)
			}
			if err.Error() == "exit+save" {
				Serialize()
				return NotifyObserver(command)
			}
			// 错误日志
			_ = NotifyObserver(nil)
			return err
		}

		err = NotifyObserver(command)
		if err != nil {
			return err
		}

	}
	return err
}

func ReadCommand(line string) (Command, error) {

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
	if commandsMapper[args[0]] == nil {
		return nil, errors.New("invalid command")
	}
	command := commandsMapper[args[0]]()
	err := command.SetArgs(args)
	if err != nil {
		return nil, err
	}
	return command, nil
}
