package workspace

import (
	"design/util"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Serialize() {
	// if backup.json is not empty, then return
	if flag, err := util.IsFileEmpty(path + "backup.json"); !flag && err == nil {
		return
	}
	if !isEmpty(CurWorkspace) {
		allWorkspaces["CurWorkspace"] = *CurWorkspace
	}

	//for _, workspace := range allWorkspaces {
	jsonData, err := json.Marshal(allWorkspaces)
	//if len(jsonData) == 0 {
	//	fmt.Println("JSON serialization error:", errors.New("empty json data"))
	//}
	if err != nil {
		fmt.Println("JSON serialization error:", err)
		return
	}

	//fmt.Println("JSON data:", string(jsonData))

	//}
	_ = util.AsJson(string(jsonData), path+"backup.json")
}

func Deserialize() {
	if flag, err := util.IsFileEmpty(path + "backup.json"); flag || err != nil {
		return
	}
	f, err := os.OpenFile(path+"backup.json", os.O_RDONLY, 0644)
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

	var workspaces = make(map[string]Workspace)
	err = json.Unmarshal([]byte(data), &workspaces)
	if err != nil {
		fmt.Println("JSON deserialization error:", err)
		return
	}
	for _, workspace := range workspaces {
		if workspace.FileName == "CurWorkspace" || workspace.FileName == "" {
			continue
		}
		allWorkspaces[workspace.FileName] = workspace
	}
	for i, workspace := range workspaces {
		if i == "CurWorkspace" && workspace.FileName != "" {
			*CurWorkspace = allWorkspaces[workspace.FileName]
			break
		}
	}
}
