package workspace

import (
	"design/util"
	"os"
	"sync"
)

func Init() {
	Deserialize()
	_ = util.AsJson("", path+"backup.json")
}

// updateWorkspace update workspace in allWorkspaces
func updateWorkspace(curWorkspace *Workspace) {
	if isEmpty(curWorkspace) {
		return
	}
	_, ok := allWorkspaces[curWorkspace.FileName]
	if ok {
		allWorkspaces[curWorkspace.FileName] = *curWorkspace
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

// Dirty if Dirty, save to file
func (curWorkspace *Workspace) Dirty() bool {
	return len(curWorkspace.UndoableCommandHistory) != 0
}

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
