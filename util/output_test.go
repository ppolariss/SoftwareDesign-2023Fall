package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	_ = Output("test", nil)
	_ = Output(" ", nil)
	_ = Output("test", nil)
	_ = Output("\n", nil)
	_ = Output("test", nil)

	filepath := "test_output"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	defer func(file *os.File) {
		_ = file.Close()
		_ = os.Remove("test_output")
	}(file)

	if err != nil {
		panic(err)
	}
	_ = Output("test", file)
	_ = Output(" ", file)
	_ = Output("test", file)
	_ = Output("\n", file)
	_ = Output("test", file)

	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	contents, err := ReadStrings(file)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []string{"test test", "test"}, contents)

	_, err = file.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	content, err := ReadString(file)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "test testtest", content)

	filepath = "test_output2"
	err = AsJson("apple pie", filepath)
	if err != nil {
		t.Fatal(err)
	}
	openFile, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = openFile.Close()
		_ = os.Remove("test_output2")
	}()
	str, err := ReadString(openFile)
	assert.Equal(t, "apple pie", str)

	filepath = "test_output3"
	err = AsFile(1, []string{"apple", "pie"}, filepath)
	if err != nil {
		t.Fatal(err)
	}
	openFile2, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = openFile2.Close()
		_ = os.Remove("test_output3")
	}()
	ss, err := ReadStrings(openFile2)
	assert.Equal(t, []string{"apple", "pie"}, ss)
}
