package workspace

import (
	"errors"
	"fmt"
)

type Exit struct {
}

func (e *Exit) Execute() error {
	return CurWorkspace.Exit()
}
func (e *Exit) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("exit: args error")
	}
	return nil
}

func (e *Exit) CallSelf() string {
	return "exit"
}

func (curWorkspace *Workspace) Exit() error {
	var dirty bool
	for _, workspace := range allWorkspaces {
		if workspace.Dirty {
			dirty = true
			break
		}
	}
	if dirty {
		fmt.Println("Do you want to save the unsaved workspace [Y\\N] ？")
		var input string
		for {
			_, err := fmt.Scanln(&input)
			if err != nil {
				return errors.New(err.Error())
			}
			if input == "Y" || input == "y" {
				for _, workspace := range allWorkspaces {
					if workspace.Dirty {
						err = workspace.Save()
						if err != nil {
							return err
						}
					}
				}
				break
			} else if input == "N" || input == "n" {
				break
			} else {
				fmt.Println("Please input Y or N")
			}
		}
	}

	allWorkspaces = make(map[string]Workspace)

	*curWorkspace = Workspace{}
	return nil
}
