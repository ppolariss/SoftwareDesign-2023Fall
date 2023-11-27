package workspace

import (
	. "design/interfaces"
)

type Workspace struct {
	FileName               string
	UndoableCommandHistory []UndoableCommand
	UndoableCommandPointer int
	FileContent            []string
	CreateAt               string
	Dirty                  bool
}

// var CurrentWorkspace *Workspace
var AllWorkspaces = make(map[string]Workspace)
var Path = "./files/"
var CurWorkspace = &Workspace{}
