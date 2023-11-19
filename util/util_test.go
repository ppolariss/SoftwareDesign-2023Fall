package util

import (
	"fmt"
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

func TestGetInteval(t *testing.T) {
	str, err := GetInterval("20201214 12:12:12", "20201213 12:12:12")
	if str != "1天" || err != nil {
		t.Error("error")
	}
	timeStr2 := "20231031 10:30:00"
	timeStr1 := "20231101 11:45:30"
	str, err = GetInterval(timeStr1, timeStr2)
	if str != "1天1小时15分钟30秒" || err != nil {
		fmt.Println(str)
		t.Error("error")
	}
}
