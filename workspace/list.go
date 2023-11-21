package workspace

import "fmt"

//	test1
//
// ->test2 *
var space = "  "
var arrow = "->"
var star = " *"

func (curWorkspace *Workspace) List() {
	for _, workspace := range allWorkspaces {
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
}
