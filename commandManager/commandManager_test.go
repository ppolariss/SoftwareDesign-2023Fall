package commandManager

import (
	"design/interfaces"
	"design/workspace"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandManager(t *testing.T) {
	var c interfaces.UndoableCommand

	workspace.CurWorkspace = &workspace.Workspace{
		FileName: "1",
		UndoableCommandHistory: []interfaces.UndoableCommand{
			c,
		},
		UndoableCommandPointer: 0,
		FileContent:            []string{},
		CreateAt:               "",
		Dirty:                  true,
	}
	p := previous()
	assert.Equal(t, nil, p)
	p = next()
	assert.Equal(t, c, p)
	assert.Equal(t, false, workspace.CurWorkspace.Dirty)
	p = previous()
	assert.Equal(t, c, p)
	assert.Equal(t, true, workspace.CurWorkspace.Dirty)
}
