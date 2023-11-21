package workspace

import (
	"errors"
	"fmt"
)

func (curWorkspace *Workspace) Close(fileName string) error {
	if isEmpty(curWorkspace) {
		return errors.New("close: curWorkspace is nil")
	}
	if curWorkspace.Dirty {
		fmt.Println("Do you want to save the current workspace [Y\\N] ？")
		var input string
		for {
			_, err := fmt.Scanln(&input)
			if err != nil {
				return errors.New(err.Error())
			}
			if input == "Y" || input == "y" {
				err = curWorkspace.Save()
				if err != nil {
					return err
				}
				break
			} else if input == "N" || input == "n" {
				break
			} else {
				fmt.Println("Please input Y or N")
			}
		}
	}

	_, ok := allWorkspaces[fileName]
	if ok {
		delete(allWorkspaces, fileName)
		//ToDo
		*curWorkspace = Workspace{}
		return nil
	}
	return errors.New("close: no such file")

}
