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
var allWorkspaces map[string]Workspace
var path string
var CurWorkspace = &Workspace{}
