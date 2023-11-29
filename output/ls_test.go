package output

import (
	"design/util"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"testing"
)

func TestLs(t *testing.T) {
	f, err := os.Create("test")
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = f
	defer func() {
		_ = f.Close()
		_ = os.Remove("test")
	}()
	err = Ls("./", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	strings, err := util.ReadStrings(f)
	if err != nil {
		return
	}
	sort.Strings(strings)
	assert.Equal(t, []string{
		"└── test",
		"├── file.go",
		"├── file_test.go",
		"├── ls.go",
		"├── ls_test.go",
		"├── output.go",
		"├── output_test.go",
	}, strings)
}
