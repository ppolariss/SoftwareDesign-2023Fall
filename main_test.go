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

func TestCommandLab2(t *testing.T) {
	err := prepare(2)
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
		err := clear()
		if err != nil {
			return
		}
		inputFile, err := os.Open("testFiles/lab2/test" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
			return
		}

		os.Stdin = inputFile

		//tmpfile, err := os.Create("testFiles/lab2/result" + strconv.Itoa(i))
		tmpfile, err := os.OpenFile("testFiles/lab2/result"+strconv.Itoa(i), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
			return
		}

		// defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile

		command.Init()
		err = command.Do()
		if err != nil && err.Error() != "EOF" {
			t.Fatal(err)
			return
		}

		_ = tmpfile.Close()
		_ = inputFile.Close()
	}

	for i := 1; i <= 5; i++ {
		filePath1 := "./testFiles/lab2/result" + strconv.Itoa(i)
		filePath2 := "./testFiles/lab2/stdresult" + strconv.Itoa(i)
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

// due to change of save function, this test is no longer valid
func TestCommandLab1(t *testing.T) {
	dirPaths := []string{
		"./files",
		"./testFiles",
		"./logFiles",
		"./testFiles/lab1",
	}
	for _, dirPath := range dirPaths {
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			// 目录不存在，创建目录
			err := os.Mkdir(dirPath, 0755) // 0755 是权限设置，可以根据需要更改
			if err != nil {
				t.Fatal()
				return
			}
		}
	}
	//_ = os.Mkdir(dirPath, os.ModePerm)
	for i := 1; i <= 5; i++ {
		filePath := "./files/test" + strconv.Itoa(i) + ".md"
		// os.Remove(file_path)
		f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
		}
		_ = f.Truncate(0)
		_ = f.Close()
	}

	//oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		//	os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	for i := 1; i <= 5; i++ {
		inputFile, err := os.Open("testFiles/lab1/test" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}

		//os.Stdin = inputFile

		//tmpfile, err := os.Create("testFiles/lab1/result" + strconv.Itoa(i))
		tmpfile, err := os.OpenFile("testFiles/lab2/result"+strconv.Itoa(i), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
		}
		err = tmpfile.Truncate(0)
		if err != nil {
			return
		}
		// defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile

		err = command.Do()
		if err != nil {
			return
		}

		_ = tmpfile.Close()
		_ = inputFile.Close()
	}

	for i := 1; i <= 5; i++ {
		filePath1 := "./testFiles/lab1/result" + strconv.Itoa(i)
		filePath2 := "./testFiles/lab1/stdresult" + strconv.Itoa(i)
		f1, err := os.Open(filePath1)
		if err != nil {
			t.Fatal(err)
		}
		f2, err := os.Open(filePath2)
		if err != nil {
			t.Fatal(err)
		}
		s1, err := util.ReadStrings(f1)
		if err != nil {
			t.Fatal(err)
		}
		s2, err := util.ReadStrings(f2)
		if err != nil {
			t.Fatal(err)
		}
		if !compareStringArrays(s1, s2) {
			t.Fatal(i, s1, s2)
		}

		// // file_path1 := "./testFiles/" + strconv.Itoa(i)
		// // file_path2 := "./testFiles/klsjfkfs hksdhfjs"
		// content1, err := os.ReadFile(file_path1)
		// if err != nil {
		// 	panic(err)
		// }

		// content2, err := os.ReadFile(file_path2)
		// if err != nil {
		// 	panic(err)
		// }

		// normalizedContent1 := normalizeLineEndings(string(content1))
		// normalizedContent2 := normalizeLineEndings(string(content2))

		// if bytes.Equal([]byte(normalizedContent1), []byte(normalizedContent2)) {
		// 	t.Fatal("\n", normalizedContent1, normalizedContent2)
		// }

		// // file1, err := os.ReadFile("./testFiles/result" + strconv.Itoa(i) )
		// // file2, _ := os.ReadFile("./testFiles/stdresult" + strconv.Itoa(i) )
		// // if err != nil {
		// // 	t.Fatal(err)
		// // }
		// // if !bytes.Equal(file1, file2) {
		// // 	// t.Fatal("testFiles" + strconv.Itoa(i) + " failed")
		// // }
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

func clear() error {
	workspace.CurWorkspace = &workspace.Workspace{}
	workspace.AllWorkspaces = make(map[string]workspace.Workspace)
	return nil
}
