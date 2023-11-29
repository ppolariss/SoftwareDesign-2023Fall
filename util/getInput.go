package util

import (
	"bufio"
	"io"
)

// var once sync.Once
var scanner *bufio.Scanner

func SetReader(r io.Reader) {
	scanner = bufio.NewScanner(r)
}

func GetInput() string {
	for {
		//scanner.Split(bufio.ScanLines)
		if scanner.Scan() {
			input := scanner.Text()
			return input
		} else {
			return ""
		}
		//if err := scanner.Err(); err != nil {
		//	return ""
		//}
	}
}
