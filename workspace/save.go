package workspace

import (
	"design/util"
	"errors"
)

type Save struct {
}

func (c *Save) Execute() error {
	//return output.AsFile(1, CurWorkspace.FileContent, GetFilePath(CurWorkspace.FileName))
	return CurWorkspace.Save()
}

func (c *Save) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("save: args error")
	}
	return nil
}

func (c *Save) CallSelf() string {
	return "save"
}

func (curWorkspace *Workspace) Save() error {
	if isEmpty(curWorkspace) {
		return errors.New("save: curWorkspace is nil")
	}
	//name := reflect.TypeOf(command).Elem().Name()
	//if name == "save" || name == "load" {
	curWorkspace.UndoableCommandHistory = curWorkspace.UndoableCommandHistory[:0]
	curWorkspace.UndoableCommandPointer = 0
	//	return nil
	//}
	updateWorkspace(curWorkspace)
	//*curWorkspace = nil
	return util.AsFile(1, curWorkspace.FileContent, GetFilePath(curWorkspace.FileName))
}
