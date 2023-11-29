package editor

import (
	"errors"
	"strconv"
	"strings"
)

// return lineNum according to content
func matchContent(content string, fileContent []string) (int, error) {
	for i, v := range fileContent {
		s := getBareContent(v)
		if s == content {
			return i + 1, nil
		}
	}
	return 0, errors.New("matchContent(): content not found")
}

func getBareContent(s string) string {
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return s
	}
	i := 0
	for {
		if s[i] == '#' {
			i++
		} else if s[i] == ' ' {
			break
		} else {
			i = 0
			break
		}
	}
	if i > 0 {
		s = s[i+1:]
	}

	ss := strings.Split(s, " ")
	if len(ss) == 1 {
		return s
	}
	// try to parse the first word to int
	if len(ss[0]) < 1 {
		return s
	}
	_, err := strconv.Atoi(ss[0][:len(ss[0])-1])
	if err != nil {
		return s
	}
	if ss[0][len(ss[0])-1] != '.' {
		return s
	}
	return strings.Join(ss[1:], " ")
}
