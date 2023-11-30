package log

import (
	"design/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logPath = "./"
	logFilePath = "./2"
	path := "./3"
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	assert.Equal(t, nil, err)
	os.Stdout = f
	defer func() {
		_ = os.Remove("./global")
		_ = os.Remove("./2")
		_ = f.Close()
		_ = os.Remove("./3")
	}()

	// test update log
	log := Log{}
	err = log.Update(nil)
	assert.Equal(t, nil, err)

	// test history
	err = history(1)
	assert.Equal(t, nil, err)
	//assert.Equal(t, "history: curWorkspace is nil", err.Error())
	_, err = f.Seek(0, 0)
	assert.Equal(t, nil, err)
	readString, err := util.ReadString(f)
	assert.Equal(t, nil, err)
	assert.Equal(t, util.GetNow()[0:13]+"error", readString[0:13]+readString[len(readString)-5:])

	err = f.Truncate(0)
	assert.Equal(t, nil, err)
	_, err = f.Seek(0, 0)
	assert.Equal(t, nil, err)

	// test stats
	err = stats("current", util.GetNow(), "1")
	_, err = f.Seek(0, 0)
	assert.Equal(t, nil, err)
	readString, err = util.ReadString(f)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1 0秒", readString)
}
