package command

import (
	. "design/interfaces"
	"design/util"
	"design/workspace"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type serialize struct {
	FileName               string
	UndoableCommandHistory []map[string]interface{}
	UndoableCommandPointer int
	FileContent            []string
	CreateAt               string
	Dirty                  bool
}

func Serialize() {
	// if backup.json is not empty, then return
	if flag, err := util.IsFileEmpty(workspace.Path + "backup.json"); !flag && err == nil {
		return
	}
	if !workspace.IsEmpty(workspace.CurWorkspace) {
		workspace.AllWorkspaces["CurWorkspace"] = *workspace.CurWorkspace
	}

	//for _, workspace := range allWorkspaces {
	jsonData, err := json.Marshal(workspace.AllWorkspaces)
	//if len(jsonData) == 0 {
	//	fmt.Println("JSON serialization error:", errors.New("empty json data"))
	//}
	if err != nil {
		fmt.Println("JSON serialization error:", err)
		return
	}

	//fmt.Println("JSON data:", string(jsonData))

	//}
	_ = util.AsJson(string(jsonData), workspace.Path+"backup.json")
}

func Deserialize() {
	if flag, err := util.IsFileEmpty(workspace.Path + "backup.json"); flag || err != nil {
		return
	}
	f, err := os.OpenFile(workspace.Path+"backup.json", os.O_RDONLY, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return
		}
		fmt.Println("Open backup.json error:", err)
		return
	}
	var data string
	data, err = util.ReadString(f)
	if err != nil {
		fmt.Println("Read backup.json error:", err)
		return
	}
	//fmt.Println("data:", data)

	var workspaces = make(map[string]serialize)
	err = json.Unmarshal([]byte(data), &workspaces)
	if err != nil {
		fmt.Println("JSON deserialization error:", err)
		return
	}
	for _, ws := range workspaces {
		if ws.FileName == "CurWorkspace" || ws.FileName == "" {
			continue
		}
		var tmp = workspace.Workspace{
			FileName:               ws.FileName,
			UndoableCommandPointer: ws.UndoableCommandPointer,
			FileContent:            ws.FileContent,
			CreateAt:               ws.CreateAt,
			Dirty:                  ws.Dirty,
		}
		tmp.UndoableCommandHistory = make([]UndoableCommand, 0)
		for _, command := range ws.UndoableCommandHistory {
			switch command["Name"] {
			case "insert":
				tmp.UndoableCommandHistory = append(tmp.UndoableCommandHistory, &InsertCommand{
					Name:    "insert",
					LineNum: (int)(command["LineNum"].(float64)),
					Content: command["Content"].(string),
				})
			case "delete":
				tmp.UndoableCommandHistory = append(tmp.UndoableCommandHistory, &DeleteCommand{
					Name:    "delete",
					LineNum: (int)(command["LineNum"].(float64)),
					Content: command["Content"].(string),
				})
			case "append-head":
				tmp.UndoableCommandHistory = append(tmp.UndoableCommandHistory, &AppendHead{
					Name:    "append-head",
					Content: command["Content"].(string),
				})
			case "append-tail":
				tmp.UndoableCommandHistory = append(tmp.UndoableCommandHistory, &AppendTail{
					Name:    "append-tail",
					LineNum: (int)(command["LineNum"].(float64)),
					Content: command["Content"].(string),
				})
			}
			workspace.AllWorkspaces[ws.FileName] = tmp
		}
		for i, ws := range workspaces {
			if i == "CurWorkspace" && ws.FileName != "" {
				*workspace.CurWorkspace = workspace.AllWorkspaces[ws.FileName]
				break
			}
		}
	}
}
