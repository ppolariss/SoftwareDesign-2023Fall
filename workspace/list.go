package workspace

import (
	"errors"
	"fmt"
)

//	test1
//
// ->test2 *
var space = "  "
var arrow = "->"
var star = " *"

type List struct {
}

func (l *List) Execute() error {
	return CurWorkspace.List()
}
func (l *List) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("list-workspace: args error")
	}
	return nil
}

func (l *List) CallSelf() string {
	return "list-workspace"
}

func (curWorkspace *Workspace) List() error {
	updateWorkspace(CurWorkspace)
	for _, workspace := range AllWorkspaces {
		if workspace.FileName == curWorkspace.FileName {
			fmt.Print(arrow, workspace.FileName)
		} else {
			fmt.Print(space, workspace.FileName)
		}
		if workspace.Dirty {
			fmt.Println(star)
		} else {
			fmt.Println()
		}
	}
	return nil
}
