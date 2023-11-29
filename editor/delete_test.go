package editor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	var (
		origin = []string{
			"1 # 1",
			"2 # 2",
			"3 # 3",
			"4 # 4",
		}
		stdResult = []string{
			"1 # 1",
			"2 # 2",
			"4 # 4",
		}
		line    int
		content = "3 # 3"
		err     error
		result  = make([]string, len(origin))
	)

	// assign content
	copy(result, origin)
	line, content, err = Delete(-1, content, &result)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, line)
	assert.Equal(t, "3 # 3", content)
	assert.Equal(t, stdResult, result)
	assert.Equal(t, 3, len(result))

	// assign lineNum
	result = make([]string, len(origin))
	copy(result, origin)
	line, content, err = Delete(3, "", &result)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, line)
	assert.Equal(t, "3 # 3", content)
	assert.Equal(t, stdResult, result)
	assert.Equal(t, 3, len(result))

	// error
	_, _, err = Delete(5, "", &result)
	assert.Equal(t, "delete: line number out of range", err.Error())

	// error
	_, _, err = Delete(-1, "apple", &result)
	assert.Equal(t, "matchContent(): content not found", err.Error())
}

// = 浅拷贝
// 直接copy不会扩容
