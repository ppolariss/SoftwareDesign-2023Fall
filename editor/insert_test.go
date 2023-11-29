package editor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
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
			"apple",
			"3 # 3",
			"4 # 4",
		}
		stdResult2 = []string{
			"1 # 1",
			"2 # 2",
			"3 # 3",
			"4 # 4",
			"apple",
		}
		line    int
		content = "apple"
		err     error
		result  []string
	)
	result = make([]string, len(origin))
	copy(result, origin)
	line, err = Insert(3, content, &result)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, line)
	assert.Equal(t, stdResult, result)
	assert.Equal(t, 5, len(result))

	result = make([]string, len(origin))
	copy(result, origin)
	line, err = Insert(-1, content, &result)
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, line)
	assert.Equal(t, stdResult2, result)
	assert.Equal(t, 5, len(result))

	result = make([]string, len(origin))
	copy(result, origin)
	line, err = Insert(5, content, &result)
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, line)
	assert.Equal(t, stdResult2, result)
	assert.Equal(t, 5, len(result))

	result = make([]string, len(origin))
	copy(result, origin)
	_, err = Insert(0, content, &result)
	assert.Equal(t, "insert: line number out of range", err.Error())
	_, err = Insert(6, content, &result)
	assert.Equal(t, "insert: line number out of range", err.Error())
	_, err = Insert(0, content, nil)
	assert.Equal(t, "insert: fileContent is nil", err.Error())
}
