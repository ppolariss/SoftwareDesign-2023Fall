package workspace

import (
	"bufio"
	"design/command"
	e "design/myError"
	"fmt"
	"os"
)

type WorkspaceInf interface {
	// SetArgs([]string) error
	// Execute() error
	// CallSelf() string
	load(fileName string) (Workspace, error)
}

type Workspace struct {
	fileName               string
	dirty                  bool // if dirty, save to file
	undoableCommandHistory []command.UndoableCommand
	undoableCommandPointer int
	fileContent            []string
}

var CurrentWorkspace Workspace
var allWorkspaces map[string]Workspace
var path string

func init() {
	allWorkspaces = make(map[string]Workspace)
	path = "./file/"
}

func updateWorkspace(curWorkspace Workspace) {
	_, ok := allWorkspaces[curWorkspace.fileName]
	if ok {
		allWorkspaces[curWorkspace.fileName] = curWorkspace
	}
	// if nil
	// else ?
}

func (curWorkspace *Workspace) Load(fileName string) (*Workspace, error) {
	if curWorkspace == nil {
		return nil, e.NewMyError("load: curWorkspace is nil")
	}
	updateWorkspace(*curWorkspace)

	ws, ok := allWorkspaces[fileName]
	if ok {
		CurrentWorkspace = ws
		return &CurrentWorkspace, nil
	}

	ws = Workspace{
		fileName:               fileName,
		dirty:                  false,
		undoableCommandHistory: make([]command.UndoableCommand, 0),
		undoableCommandPointer: 0,
		fileContent:            make([]string, 0),
	}

	filePath := path + fileName

	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, e.NewMyError(err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				return nil, e.NewMyError(err.Error())
			}

			if content == "" {
				break
			} else {
				ws.fileContent = append(ws.fileContent, content)
				break
			}
			// 注意此处还要处理最后一行
		}
		// if content == "" {
		// 	continue
		// }
		ws.fileContent = append(ws.fileContent, content)
	}
	CurrentWorkspace = ws
	return &CurrentWorkspace, nil
}

func (curWorkspace *Workspace) Open(fileName string) (*Workspace, error) {
	if curWorkspace == nil {
		return nil, e.NewMyError("open: curWorkspace is nil")
	}
	updateWorkspace(*curWorkspace)
	_, ok := allWorkspaces[fileName]
	if ok {
		CurrentWorkspace = allWorkspaces[fileName]
		return &CurrentWorkspace, nil
	}
	return nil, e.NewMyError("open: no such file")
}

func (curWorkspace *Workspace) Close(fileName string) error {
	if curWorkspace == nil {
		return e.NewMyError("close: curWorkspace is nil")
	}
	if curWorkspace.dirty {
		fmt.Println("Do you want to save the current workspace [Y\\N] ？")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return e.NewMyError(err.Error())
		}
		if input == "Y" || input == "y" {
			updateWorkspace(*curWorkspace)
		}
	}

	_, ok := allWorkspaces[fileName]
	if ok {
		delete(allWorkspaces, fileName)
		return nil
	}
	return e.NewMyError("close: no such file")

}
