package workspace

import (
	. "design/interfaces"
)

type Workspace struct {
	FileName               string
	Dirty                  bool // if Dirty, save to file
	UndoableCommandHistory []UndoableCommand
	UndoableCommandPointer int
	FileContent            []string
	CreateAt               string
}

// var CurrentWorkspace *Workspace
var allWorkspaces map[string]Workspace
var path string
