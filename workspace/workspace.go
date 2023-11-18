package workspace

import (
	"bufio"
	"design/command"
	e "design/myError"
	"os"
)

type Workspace struct {
	fileName               string
	dirty                  bool
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

func load(fileName string, workspace Workspace) (Workspace, error) {
	_, ok := allWorkspaces[workspace.fileName]
	if ok {
		allWorkspaces[workspace.fileName] = workspace
	}
	// if nil
	// else ?

	ws, ok := allWorkspaces[fileName]
	if ok {
		CurrentWorkspace = ws
		return ws, nil
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
		return ws, e.NewMyError(err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				return ws, e.NewMyError(err.Error())
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
	return ws, nil
}
