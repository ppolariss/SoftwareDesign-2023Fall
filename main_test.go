package main

import (
	"design/command"
	"design/workspace"
	"fmt"
	"strings"

	// "bytes"
	// "fmt"
	"design/util"
	"os"
	"strconv"
	"testing"
)

func Test2(t *testing.T) {
	testCommand(2, t)
}

// due to change of save function, this test is no longer valid
//func Test1(t *testing.T) {
//	testCommand(1, t)
//}

func testCommand(nLab int, t *testing.T) {
	err := prepare(nLab)
	if err != nil {
		t.Fatal(err)
		return
	}

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()
	for i := 1; i <= 5; i++ {
		err := clearSpace()
		if err != nil {
			return
		}
		inputFile, err := os.Open("testFiles/lab" + strconv.Itoa(nLab) + "/test" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
			return
		}

		os.Stdin = inputFile

		//tmpfile, err := os.Create("testFiles/lab2/result" + strconv.Itoa(i))
		tmpfile, err := os.OpenFile("testFiles/lab"+strconv.Itoa(nLab)+"/result"+strconv.Itoa(i), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
			return
		}

		// defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile

		command.Init(inputFile)
		err = command.Do()
		if err != nil && err.Error() != "EOF" {
			t.Fatal(err)
			return
		}

		_ = tmpfile.Close()
		_ = inputFile.Close()
	}

	for i := 1; i <= 5; i++ {
		filePath1 := "./testFiles/lab" + strconv.Itoa(nLab) + "/result" + strconv.Itoa(i)
		filePath2 := "./testFiles/lab" + strconv.Itoa(nLab) + "/stdresult" + strconv.Itoa(i)
		f1, err := os.Open(filePath1)
		if err != nil {
			t.Fatal(err)
			return
		}
		f2, err := os.Open(filePath2)
		if err != nil {
			t.Fatal(err)
			return
		}
		s1, err := util.ReadStrings(f1)
		if err != nil {
			t.Fatal(err)
			return
		}
		s2, err := util.ReadStrings(f2)
		if err != nil {
			t.Fatal(err)
			return
		}
		if !compareStringArrays(s1, s2) {
			fmt.Println(i, s1, s2)
			t.Fatal(i, s1, s2)
			return
		}
	}
}

// std arr2
func compareStringArrays(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] == "" && i == len(arr1)-1 {
			continue
		}
		if arr2[i] == "***" {
			continue
		}
		if strings.HasSuffix(arr1[i], "\r") {
			arr1[i] = arr1[i][:len(arr1[i])-1]
		}
		if strings.HasSuffix(arr2[i], "\r") {
			arr2[i] = arr2[i][:len(arr2[i])-1]
		}
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func prepare(index int) error {
	dirPaths := []string{
		"./files",
		"./testFiles",
		"./logFiles",
		"./testFiles/lab" + strconv.Itoa(index),
	}
	for _, dirPath := range dirPaths {
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// 目录不存在，创建目录
			err := os.Mkdir(dirPath, 0755) // 0755 是权限设置，可以根据需要更改
			if err != nil {
				return err
			}
		}
	}
	//_ = os.Mkdir(dirPath, os.ModePerm)
	for i := 1; i <= 7; i++ {
		filePath := "./files/test" + strconv.Itoa(i) + ".md"
		_ = os.Remove(filePath)
		//f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		//if err != nil {
		//	return err
		//}
		//_ = f.Truncate(0)
		//_ = f.Close()
	}

	f, err := os.OpenFile("./files/backup.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	_ = f.Truncate(0)
	_ = f.Close()

	for i := 1; i <= 5; i++ {
		filePath := "./testFiles/lab" + strconv.Itoa(index) + "/result" + strconv.Itoa(i)
		f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		_ = f.Truncate(0)
		_ = f.Close()
	}
	return nil
}

func clearSpace() error {
	workspace.CurWorkspace = &workspace.Workspace{}
	workspace.AllWorkspaces = make(map[string]workspace.Workspace)
	return nil
}
