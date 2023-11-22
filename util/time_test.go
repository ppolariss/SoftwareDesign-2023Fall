package util

import (
	"fmt"
	"testing"
)

func TestGetInterval(t *testing.T) {
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
