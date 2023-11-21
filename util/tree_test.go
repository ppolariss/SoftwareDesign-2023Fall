package util

//
//import (
//	"fmt"
//	"os"
//	"testing"
//)
//
//func TestParseFromFile(t *testing.T) {
//	// 相对位置 取决于运行所在
//	filePath := "../file/test_parse.txt"
//	// os.Create(filepath)
//	// file, _ := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
//	// file.WriteString("# 我的资源\n## 程序设计\n### 软件设计\n#### 设计模式\n1. 观察者模式\n## 工具箱\n### Adobe")
//	// file.Close()
//
//	err := parseFromFile(filePath)
//	defer func(name string) {
//		_ = os.Remove(name)
//	}(filePath)
//	if err != nil {
//		panic(err)
//	}
//	root := GetRoot()
//	if root.children[0].content != "我的资源" {
//		fmt.Println(root.children[0].content)
//		t.Errorf("parseFromFile() failed")
//	}
//	if root.children[0].children[0].content != "程序设计" {
//		fmt.Println(root.children[0].children[0].content)
//		t.Errorf("parseFromFile() failed")
//	}
//	if root.children[0].children[0].children[0].content != "软件设计" {
//		fmt.Println(len(root.children[0].children[0].children[0].content))
//		t.Errorf("parseFromFile() failed")
//	}
//	if root.children[0].children[0].children[0].children[0].content != "设计模式" {
//		fmt.Println(root.children[0].children[0].children[0].children[0].content)
//		t.Errorf("parseFromFile() failed")
//	}
//	if root.children[0].children[0].children[0].children[0].children[0].content != "1. 观察者模式" {
//		fmt.Println(root.children[0].children[0].children[0].children[0].children[0].content)
//		t.Errorf("parseFromFile() failed")
//	}
//	if root.children[0].children[1].content != "工具箱" {
//		fmt.Println(len(root.children[0].children[1].content))
//		t.Errorf("parseFromFile() failed")
//	}
//	if root.children[0].children[1].children[0].content != "Adobe" {
//		fmt.Println(root.children[0].children[1].children[0].content)
//		t.Errorf("parseFromFile() failed")
//	}
//	// Dump()
//}
