package util

import (
	"bufio"
	"io"
)

// var once sync.Once
var scanner *bufio.Scanner

func GetInput(r io.Reader) string {
	once.Do(func() {
		scanner = bufio.NewScanner(r)

	})
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
