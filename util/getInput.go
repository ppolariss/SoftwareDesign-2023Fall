package util

import (
	"bufio"
	"io"
)

var scanner *bufio.Scanner

func SetReader(r io.Reader) {
	scanner = bufio.NewScanner(r)
}

// GetInput Due to the need to retrieve content from standard input at multiple points
// to facilitate reading from a file during testing,
// the input has been encapsulated into GetInput function
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

// you can't use fmt.Fscanln or fmt.Fscan to read
// because it will throw error when encounter space and newline
// and if you ignore the error, it will eat off the next letter
