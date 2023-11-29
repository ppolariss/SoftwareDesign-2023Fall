package util

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInput(t *testing.T) {
	var input string
	SetReader(nil)
	input = GetInput()
	assert.Equal(t, "", input)

	data := []byte("test input")
	SetReader(bytes.NewReader(data))
	input = GetInput()
	assert.Equal(t, "test input", input)
	data = []byte("test input\n input")
	SetReader(bytes.NewReader(data))
	input = GetInput()
	assert.Equal(t, "test input", input)
	input = GetInput()
	assert.Equal(t, " input", input)
}
