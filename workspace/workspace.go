package workspace

import (
	"design/util"
	"os"
	"sync"
)

func Init() {
}

// updateWorkspace update workspace in AllWorkspaces
func updateWorkspace(curWorkspace *Workspace) {
	if IsEmpty(curWorkspace) {
		return
	}
	_, ok := AllWorkspaces[curWorkspace.FileName]
	if ok {
		AllWorkspaces[curWorkspace.FileName] = *curWorkspace
	}
}

func IsEmpty(workspace *Workspace) bool {
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

//// Dirty if Dirty, save to file
//func (curWorkspace *Workspace) Dirty() bool {
//	return len(curWorkspace.UndoableCommandHistory) != 0
//}

var once sync.Once

// Log workspace for stats
func Log(curWorkspace *Workspace) error {
	if curWorkspace == nil {
		return nil
	}

	interval, err := util.GetInterval(util.GetNow(), curWorkspace.CreateAt)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("./logFiles/logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	once.Do(func() {
		_ = f.Truncate(0)
		_ = util.Output("session start at "+curWorkspace.CreateAt+"\n", f)
	})
	return util.Output(curWorkspace.FileName+" "+interval+"\n", f)
}
