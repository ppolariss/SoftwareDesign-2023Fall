package command

import (
	"design/interfaces"
	"design/util"
	"design/workspace"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSerialize(t *testing.T) {
	var ws = workspace.Workspace{
		FileName:               "1",
		UndoableCommandHistory: []interfaces.UndoableCommand{},
		UndoableCommandPointer: 0,
		FileContent: []string{
			"ad sd",
		},
		CreateAt: "",
		Dirty:    false,
	}
	f, err := os.OpenFile("./testbackup.json", os.O_CREATE, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove("./testbackup.json")
	}()

	workspace.Path = "./test"
	workspace.AllWorkspaces["1"] = ws
	workspace.CurWorkspace = &ws

	Serialize()

	workspace.CurWorkspace = &workspace.Workspace{}
	workspace.AllWorkspaces = make(map[string]workspace.Workspace)

	Deserialize()

	_ = util.AsJson("", workspace.Path+"backup.json")
	assert.Equal(t, ws, *workspace.CurWorkspace)
	assert.Equal(t, 1, len(workspace.AllWorkspaces))
}
