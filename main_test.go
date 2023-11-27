package main

import (
	"design/command"
	// "bytes"
	// "fmt"
	"design/util"
	"os"
	"strconv"
	"testing"
)

func TestCommandLab2(t *testing.T) {
	dirPaths := []string{
		"./files",
		"./testFiles",
		"./logFiles",
		"./testFiles/lab2",
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
	for i := 1; i <= 1; i++ {
		inputFile, err := os.Open("testFiles/lab1/test" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}

		os.Stdin = inputFile

		tmpfile, err := os.Create("testFiles/lab1/result" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		// defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile

		err = command.Do(inputFile)
		if err != nil {
			return
		}

		_ = tmpfile.Close()
		_ = inputFile.Close()
	}

	for i := 1; i <= 1; i++ {
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

		os.Stdin = inputFile

		tmpfile, err := os.Create("testFiles/lab1/result" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		// defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile

		err = command.Do(inputFile)
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
		if arr2[i] == "***" {
			continue
		}
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
