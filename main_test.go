package main

import (
	// "bytes"
	// "fmt"
	"design/util"
	"os"
	"strconv"
	"testing"
)

func TestCommand(t *testing.T) {
	for i := 1; i <= 5; i++ {
		filePath := "./file/test" + strconv.Itoa(i) + ".md"
		// os.Remove(file_path)
		f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
		}
		_ = f.Truncate(0)
		_ = f.Close()
	}

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {

		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	for i := 1; i <= 5; i++ {
		inputFile, err := os.Open("test/test" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}

		os.Stdin = inputFile

		tmpfile, err := os.Create("test/result" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		// defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile

		main()

		_ = tmpfile.Close()
		_ = inputFile.Close()
	}

	for i := 1; i <= 5; i++ {
		filePath1 := "./test/result" + strconv.Itoa(i)
		filePath2 := "./test/stdresult" + strconv.Itoa(i)
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

		// // file_path1 := "./test/" + strconv.Itoa(i)
		// // file_path2 := "./test/klsjfkfs hksdhfjs"
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

		// // file1, err := os.ReadFile("./test/result" + strconv.Itoa(i) )
		// // file2, _ := os.ReadFile("./test/stdresult" + strconv.Itoa(i) )
		// // if err != nil {
		// // 	t.Fatal(err)
		// // }
		// // if !bytes.Equal(file1, file2) {
		// // 	// t.Fatal("test" + strconv.Itoa(i) + " failed")
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
