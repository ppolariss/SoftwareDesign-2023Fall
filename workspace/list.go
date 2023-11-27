package workspace

import (
	"errors"
	"fmt"
	"sort"
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
	// output as directory order
	keys := make([]string, 0, len(AllWorkspaces))
	for k := range AllWorkspaces {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {

		if AllWorkspaces[k].FileName == curWorkspace.FileName {
			fmt.Print(arrow, AllWorkspaces[k].FileName)
		} else {
			fmt.Print(space, AllWorkspaces[k].FileName)
		}
		if AllWorkspaces[k].Dirty {
			fmt.Println(star)
		} else {
			fmt.Println()
		}
	}
	return nil
}
