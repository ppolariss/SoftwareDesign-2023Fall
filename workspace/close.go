package workspace

import (
	"errors"
	"fmt"
)

type Close struct {
}

func (c *Close) Execute() error {
	return CurWorkspace.Close()
}
func (c *Close) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("close: args error")
	}
	return nil
}

func (c *Close) CallSelf() string {
	return "close"
}

func (curWorkspace *Workspace) Close() error {
	if IsEmpty(curWorkspace) {
		return errors.New("close: curWorkspace is nil")
	}
	if curWorkspace.Dirty {
		fmt.Println("Do you want to save the current workspace [Y\\N] ？")
		var input string
		for {
			_, err := fmt.Scanln(&input)
			if err != nil {
				//if err.Error() == "EOF" {
				//	continue
				//}
				return err
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

	_, ok := AllWorkspaces[curWorkspace.FileName]
	if ok {
		err := Log(curWorkspace)
		if err != nil {
			return err
		}
		delete(AllWorkspaces, curWorkspace.FileName)
		*curWorkspace = Workspace{}
		return nil
	}
	return errors.New("close: no such file")

}
