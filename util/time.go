package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func GetInterval(nowTime string, createTime string) (string, error) {
	if nowTime == "" || createTime == "" {
		return "", errors.New("GetInterval: nowTime or createTime is empty")
	}
	nowTime = strings.TrimRight(nowTime, "\n")
	createTime = strings.TrimRight(createTime, "\n")
	layout := "20060102 15:04:05"
	now, err := time.Parse(layout, nowTime)
	if err != nil {
		return "", errors.New(err.Error())
	}
	create, err := time.Parse(layout, createTime)
	if err != nil {
		return "", errors.New(err.Error())
	}

	duration := now.Sub(create)
	hours := int(duration.Hours())
	days := hours / 24
	hours = hours % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	retStr := ""
	if days != 0 {
		retStr += fmt.Sprintf("%d天", days)
	}
	if hours != 0 {
		retStr += fmt.Sprintf("%d小时", hours)
	}
	if minutes != 0 {
		retStr += fmt.Sprintf("%d分钟", minutes)
	}
	if seconds != 0 {
		retStr += fmt.Sprintf("%d秒", seconds)
	}
	if retStr == "" {
		retStr = "0秒"
	}
	return retStr, nil
}

const formatStr = "20060102 15:04:05"

func GetNow() string {
	return time.Now().Format(formatStr)
}
