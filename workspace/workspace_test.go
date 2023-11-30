package workspace

import (
	. "design/interfaces"
	"design/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWorkspace(t *testing.T) {
	var ws = Workspace{
		FileName:               "1",
		UndoableCommandHistory: []UndoableCommand{},
		UndoableCommandPointer: 0,
		FileContent: []string{
			"1",
		},
		CreateAt: util.GetNow(),
		Dirty:    false,
	}
	AllWorkspaces["1"] = ws
	*CurWorkspace = ws
	Path = "./"

	// test init
	assert.Equal(t, ws, *CurWorkspace)
	file, err := os.OpenFile("./2", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
		_ = os.Remove("./2")
	}()

	// test load opening file
	err = CurWorkspace.Load("1")
	assert.Equal(t, nil, err)
	assert.Equal(t, ws, *CurWorkspace)

	// test open non-existing file
	err = CurWorkspace.Open("2")
	assert.Equal(t, "open: no such file", err.Error())

	// test load file
	err = CurWorkspace.Load("2")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, ws, *CurWorkspace)

	// test close
	err = CurWorkspace.Close()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, IsEmpty(CurWorkspace))

	// test open
	err = CurWorkspace.Open("1")
	assert.Equal(t, nil, err)
	assert.Equal(t, ws, *CurWorkspace)

	// test list-workspaces
	_, err = file.Seek(0, 0)
	assert.Equal(t, nil, err)
	os.Stdout = file
	err = CurWorkspace.List()
	assert.Equal(t, nil, err)
	_, err = file.Seek(0, 0)
	assert.Equal(t, nil, err)
	s, err := util.ReadString(file)
	assert.Equal(t, nil, err)
	assert.Equal(t, "->1", s)

	// test save
	err = CurWorkspace.Save()
	assert.Equal(t, nil, err)
	err = os.Remove("./1")
	assert.Equal(t, nil, err)

	// test list-workspaces
	err = file.Truncate(0)
	assert.Equal(t, nil, err)
	_, err = file.Seek(0, 0)
	assert.Equal(t, nil, err)
	err = CurWorkspace.List()
	assert.Equal(t, nil, err)
	_, err = file.Seek(0, 0)
	assert.Equal(t, nil, err)
	s, err = util.ReadString(file)
	assert.Equal(t, nil, err)
	assert.Equal(t, "->1", s)
}
