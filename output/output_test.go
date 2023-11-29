package output

import (
	"design/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAsDir(t *testing.T) {
	err := AsDir("apple", []string{"banana", "cherry"})
	assert.Equal(t, "OutputAsDir(): No such node", err.Error())

	err = AsDir("apple", nil)
	assert.Equal(t, "OutputAsDir(): No file in workspace", err.Error())

	file, err := os.Create("test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = file.Close()
		_ = os.Remove("test")
	}()
	os.Stdout = file
	err = AsDir("banana", []string{"# banana", "## cherry"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	strings, err := util.ReadStrings(file)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []string{
		"└── banana",
		"    └── cherry",
	}, strings)
}
