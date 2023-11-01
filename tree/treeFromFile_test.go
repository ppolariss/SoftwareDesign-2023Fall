package tree

import (
	"fmt"
	// "os"
	// "text/template/parse"
	// "design/tree"
	// "path/filepath"
	"testing"
)

func TestParseFromFile(t *testing.T) {
	file_path := "../file/test.txt"//为什么检测不到
	// os.Create(filepath)
	// file, _ := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	// file.WriteString("# 我的资源\n## 程序设计\n### 软件设计\n#### 设计模式\n1. 观察者模式\n## 工具箱\n### Adobe")
	// file.Close()

	_, err:= parseFromFile(file_path)
	if err != nil {
		panic(err)
	}
	root := GetRoot()
	if root.children[0].content != "我的资源" {
		fmt.Println(root.children[0].content)
		// t.Errorf("parseFromFile() failed")
	}
	if root.children[0].children[0].content != "程序设计" {
		fmt.Println(root.children[0].children[0].content)
		// t.Errorf("parseFromFile() failed")
	}
	if root.children[0].children[0].children[0].content != "软件设计" {
		fmt.Println(len(root.children[0].children[0].children[0].content))
		// t.Errorf("parseFromFile() failed")
	}
	if root.children[0].children[0].children[0].children[0].content != "设计模式" {
		fmt.Println(root.children[0].children[0].children[0].children[0].content)
		// t.Errorf("parseFromFile() failed")
	}
	if root.children[0].children[0].children[0].children[0].children[0].content != "1. 观察者模式" {
		fmt.Println(root.children[0].children[0].children[0].children[0].children[0].content)
		// t.Errorf("parseFromFile() failed")
	}
	if root.children[0].children[1].content != "工具箱" {
		fmt.Println(len(root.children[0].children[1].content))
		// t.Errorf("parseFromFile() failed")
	}
	if root.children[0].children[1].children[0].content != "Adobe" {
		fmt.Println(root.children[0].children[1].children[0].content)
		t.Errorf("parseFromFile() failed")
	}
	// Dump()

	// defer os.Remove(filepath)
}