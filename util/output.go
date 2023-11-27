package util

import (
	"errors"
	"os"
)

func Output(content string, file *os.File) error {
	if file == nil {
		file = os.Stdout
	}
	_, err := file.WriteString(content)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// OutputAsFile para: 0: output to terminal; 1: output to file
func AsFile(para int, fileContent []string, filePath string) error {
	var file *os.File
	var err error
	if para == 1 {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return errors.New(err.Error())
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		// 清空文件内容
		err = file.Truncate(0)
		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		file = os.Stdout
	}
	for _, line := range fileContent {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return errors.New(err.Error())
		}
	}

	//err := util.String2tree(fileContent)
	//if err != nil {
	//	return err
	//}
	//if !util.IsInit() {
	//	return errors.New("OutputAsFile(): No file in workspace")
	//}
	//
	//tree := util.GetRoot()
	//if para != 0 {
	//	// 向文件输入
	//
	//	for _, child := range tree.Children {
	//		recurOutputAsFile(child, file)
	//	}
	//
	//} else {
	//	// 向终端输入
	//	for _, child := range tree.Children {
	//		recurOutputAsFile(child, nil)
	//	}
	//}
	return nil
}

func AsJson(content string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	// 清空文件内容
	err = file.Truncate(0)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = file.WriteString(content)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

//func recurOutputAsFile(node *util.Node, file *os.File) {
//	if node.Grade != 0 {
//		for i := 0; i < node.Grade; i++ {
//			_ = util.Output("#", file)
//
//		}
//		_ = util.Output(" ", file)
//
//	}
//
//	_ = util.Output(node.Content+"\n", file)
//	for _, child := range node.Children {
//		recurOutputAsFile(child, file)
//	}
//}
