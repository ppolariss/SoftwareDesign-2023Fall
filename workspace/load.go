package workspace

import (
	"bytes"
	. "design/interfaces"
	"design/util"
	"errors"
	"os"
)

func (curWorkspace *Workspace) Load(fileName string) error {
	ws, ok := allWorkspaces[fileName]
	if ok {
		return errors.New("load: file already opened")
		//*CurrentWorkspace = ws
		//return nil
	}

	if !isEmpty(curWorkspace) {
		updateWorkspace(curWorkspace)
	}

	ws = Workspace{
		FileName:               fileName,
		Dirty:                  false,
		UndoableCommandHistory: make([]UndoableCommand, 0),
		UndoableCommandPointer: 0,
		FileContent:            make([]string, 0),
	}

	filePath := GetFilePath(fileName)

	_, err := os.Stat(filePath)

	if err == nil {
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		lines := bytes.Split(fileBytes, []byte("\n"))
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			ws.FileContent = append(ws.FileContent, string(line))
		}
	} else if os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	ws.CreateAt = util.GetNow()
	*curWorkspace = ws
	return nil
}
