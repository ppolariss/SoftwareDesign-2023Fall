package workspace

import (
	"bytes"
	. "design/interfaces"
	"design/util"
	"errors"
	"os"
)

type Load struct {
	filepath string
}

func (c *Load) Execute() error {
	// filepath := "../file/testFiles.txt"
	// 通过main.go运行，相对路径名要从main.go所在的目录开始！！！

	return CurWorkspace.Load(c.filepath)
}

func (c *Load) SetArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("load: args error")
	}
	c.filepath = args[1]
	return nil
}

func (c *Load) CallSelf() string {
	return "load " + c.filepath
}

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
