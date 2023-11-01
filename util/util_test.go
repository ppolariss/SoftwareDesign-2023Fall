package util

import (
	"os"
	// "path/filepath"
	"testing"
)

func TestOutput(t *testing.T) {
	Output("test", nil)
	Output(" ", nil)
	Output("test", nil)
	Output("\n", nil)
	Output("test", nil)

	filepath := "../file/test_output.txt"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	Output("test", file)
	Output(" ", file)
	Output("test", file)
	Output("\n", file)
	Output("test", file)
	
	defer file.Close()
}
