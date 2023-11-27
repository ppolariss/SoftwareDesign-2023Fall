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
}

// var CurrentWorkspace *Workspace
var allWorkspaces = make(map[string]Workspace)
var path = "./files/"
var CurWorkspace = &Workspace{}
