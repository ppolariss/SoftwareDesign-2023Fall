package util

import (
	// "bufio"
	e "design/myError"
	"fmt"
	"os"
	// "github.com/go-delve/delve/pkg/dwarf/reader"
)

func Output(content string, file *os.File) error {
	if file == nil {
		fmt.Print(content)
	} else {
		_, err := file.WriteString(content)
		if err != nil {
			return e.NewMyError(err.Error())
		}
	}
	return nil
}

// func tree() {
// 	// ├── 新的标题
// 	// └── 我的书签
// 	// 	└── 学习资源
// 	// 		├── 新的⼦标题
// 	// 		└── 编程
// 	// 			└── ·新的⽂本
// }

// func main() {
// 	write()

// 	var path string = "test.txt"
// 	file, err := os.Open(path)
// 	defer file.Close()
// 	if err != nil {
// 		file, _ = os.Create(path)
// 	}
// 	// bufio ioutil
// 	reader := bufio.NewReader(file)
// 	content, err := reader.ReadString('\n')
// 	if err != nil {
// 		// panic(err)
// 		fmt.Println(err)
// 	}
// 	fmt.Println(content)

// }

// // 时间转化
// // 树结构输出


// func write() {
// 	file_path := "test.txt"

// 	file, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 		// log.Fatal(err)
// 	}
// 	defer file.Close()

// 	if _, err := file.WriteString("This is some extra text."); err != nil {
// 		panic(err)
// 	}
// }