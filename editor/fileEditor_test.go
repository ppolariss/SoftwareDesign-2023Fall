package editor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchContent(t *testing.T) {
	var (
		origin = []string{
			"1 # 1",
			"2 # 2",
			"3 # 3",
			"4 # 4",
		}
		content = "3 # 3"
		result  int
		err     error
	)
	result, err = matchContent(content, origin)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, result)
}

func TestGetBareContent(t *testing.T) {
	var (
		origin = []string{
			"1 # 1",
			"#2",
			"# 3",
			" ",
			"a",
			"1. a",
			"1. ",
			"a. a",
		}
		result string
	)
	result = getBareContent(origin[0])
	assert.Equal(t, "1 # 1", result)
	result = getBareContent(origin[1])
	assert.Equal(t, "#2", result)
	result = getBareContent(origin[2])
	assert.Equal(t, "3", result)
	result = getBareContent(origin[3])
	assert.Equal(t, " ", result)
	result = getBareContent(origin[4])
	assert.Equal(t, "a", result)
	result = getBareContent(origin[5])
	assert.Equal(t, "a", result)
	result = getBareContent(origin[6])
	assert.Equal(t, "", result)
	result = getBareContent(origin[7])
	assert.Equal(t, "a. a", result)
}
