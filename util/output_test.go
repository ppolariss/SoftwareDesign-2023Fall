package util

import (
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	_ = Output("test", nil)
	_ = Output(" ", nil)
	_ = Output("test", nil)
	_ = Output("\n", nil)
	_ = Output("test", nil)

	filepath := "../files/test_output.txt"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	_ = Output("test", file)
	_ = Output(" ", file)
	_ = Output("test", file)
	_ = Output("\n", file)
	_ = Output("test", file)

	defer func(file *os.File) {
		_ = file.Close()
	}(file)
}
