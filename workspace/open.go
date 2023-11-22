package workspace

import "errors"

type Open struct {
	fileName string
}

func (o *Open) Execute() error {
	return CurWorkspace.Open(o.fileName)
}
func (o *Open) SetArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("open: args error")
	}
	o.fileName = args[1]
	return nil
}

func (o *Open) CallSelf() string {
	return "open " + o.fileName
}

func (curWorkspace *Workspace) Open(fileName string) error {
	if curWorkspace.FileName == fileName {
		return nil
	}
	if !isEmpty(curWorkspace) {
		updateWorkspace(curWorkspace)
		err := Log(curWorkspace)
		if err != nil {
			return err
		}
	}

	_, ok := allWorkspaces[fileName]
	if ok {
		*curWorkspace = allWorkspaces[fileName]
		return nil
	}
	return errors.New("open: no such file")
}
