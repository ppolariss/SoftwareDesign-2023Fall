package workspace

import (
	"design/output"
	"errors"
)

func (curWorkspace *Workspace) Save() error {
	if isEmpty(curWorkspace) {
		return errors.New("save: curWorkspace is nil")
	}
	updateWorkspace(curWorkspace)
	//*curWorkspace = nil
	return output.OutputAsFile(1, curWorkspace.FileContent, GetFilePath(curWorkspace.FileName))
}
