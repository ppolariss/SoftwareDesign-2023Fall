package command

import (
	"design/workspace"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCommand(t *testing.T) {
	workspace.CurWorkspace = &workspace.Workspace{
		FileName:               "",
		UndoableCommandHistory: nil,
		UndoableCommandPointer: 0,
		FileContent: []string{
			"a",
			"a",
			"a",
			"a",
			"a",
		},
		CreateAt: "",
		Dirty:    false,
	}
	command, err := ReadCommand("insert 3 # haha")
	if err != nil {
		t.Fatal(err)
	}
	var args = []string{
		"insert",
		"4",
		"##",
		"hhh",
	}
	err = command.SetArgs(args)
	if err != nil {
		t.Fatal()
	}
	insertCommand := command.(*InsertCommand)
	assert.Equal(t, 4, insertCommand.LineNum)
	assert.Equal(t, "## hhh", insertCommand.Content)
}
