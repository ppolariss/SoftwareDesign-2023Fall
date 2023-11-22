package workspace

import (
	"strings"
)

func init() {
	allWorkspaces = make(map[string]Workspace)
	path = "./files/"
}

func updateWorkspace(curWorkspace *Workspace) {
	if isEmpty(curWorkspace) {
		return
	}
	_, ok := allWorkspaces[curWorkspace.FileName]
	if ok {
		allWorkspaces[curWorkspace.FileName] = *curWorkspace
	}
}

func GetFilePath(fileName string) string {
	if !strings.HasPrefix(fileName, path) {
		return path + fileName
	} else {
		return fileName
	}
}

func isEmpty(workspace *Workspace) bool {
	if workspace == nil {
		return true
	}
	if workspace.FileContent == nil || len(workspace.FileContent) == 0 {
		return true
	}
	if workspace.FileName == "" {
		return true
	}
	return false
}

// if Dirty, save to file
func (curWorkspace *Workspace) Dirty() bool {
	return len(curWorkspace.UndoableCommandHistory) != 0
}
