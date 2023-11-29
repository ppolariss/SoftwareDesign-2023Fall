package output

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestFile(t *testing.T) {
	file := &File{path: "./", name: "output"}
	if file.Name() != "output" {
		t.Fatal("file.Name() != \"output\"")
	}
	if file.GetChildren() == nil {
		t.Fatal("file.GetChildren() == nil")
	}
	if len(file.GetChildren()) == 0 {
		t.Fatal("len(file.GetChildren()) == 0")
	}
	if file.GetChildren()[0].Name() == "" {
		t.Fatal("file.GetChildren()[0].Name() == \"\"")
	}
	if file.GetChildren()[0].GetChildren() != nil {
		t.Fatal("file.GetChildren()[0].GetChildren() != nil")
	}

	assert.Equal(t, "output", file.Name())
	var s []string
	for _, child := range file.GetChildren() {
		s = append(s, child.Name())
	}
	sort.Strings(s)
	assert.Equal(t, []string{
		"file.go",
		"file_test.go",
		"ls.go",
		"ls_test.go",
		"output.go",
		"output_test.go",
	}, s)
}
