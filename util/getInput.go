// // package util
// //
// // import (
// //
// //	"bufio"
// //	"os"
// //
// // )
// //
// //var r = os.Stdin
// //
// //	func GetInput() (string, error) {
// //		scanner := bufio.NewScanner(r)
// //		scanner.Split(bufio.ScanLines)
// //		return scanner.Text(), nil
// //		//line := scanner.Text()
// //		//if err := scanner.Err(); err != nil {
// //		//	return nil, errors.New("input error")
// //		//}
// //	}
//package util
//
//import (
//	"bufio"
//	"os"
//)
//
//var r = os.Stdin
//
//func GetInput() <-chan string {
//	out := make(chan string)
//
//	go func() {
//		defer close(out)
//
//		scanner := bufio.NewScanner(r)
//		for scanner.Scan() {
//			out <- scanner.Text()
//		}
//
//		if err := scanner.Err(); err != nil {
//			// 发生错误时，将错误信息发送到通道
//			out <- err.Error()
//		}
//	}()
//
//	return out
//}

package util

import (
	"bufio"
	"errors"
	"os"
	"sync"
)

var (
	r        = os.Stdin
	input    string
	inputMux sync.Mutex
)

func GetInput() (string, error) {
	inputMux.Lock()
	defer inputMux.Unlock()

	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		input = scanner.Text()
		return input, nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New("End of input")
}
