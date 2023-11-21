package workspace

import "errors"

func (curWorkspace *Workspace) Open(fileName string) error {
	if !isEmpty(curWorkspace) {
		updateWorkspace(curWorkspace)
	}

	_, ok := allWorkspaces[fileName]
	if ok {
		*curWorkspace = allWorkspaces[fileName]
		return nil
	}
	return errors.New("open: no such file")
}
