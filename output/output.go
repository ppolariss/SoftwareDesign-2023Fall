package output

import (
	"errors"
	"fmt"
	"os"

	. "design/interfaces"
	"design/util"
)

const endBranch = "└── "
const branch = "├── "
const space = "    "
const notSpace = "│   "

func recurOutputAsFile(node *util.Node, file *os.File) {
	if node.Grade != 0 {
		for i := 0; i < node.Grade; i++ {
			_ = util.Output("#", file)

		}
		_ = util.Output(" ", file)

	}

	_ = util.Output(node.Content+"\n", file)
	for _, child := range node.Children {
		recurOutputAsFile(child, file)
	}
}

// OutputAsFile para: 0: output to terminal; 1: output to file
func OutputAsFile(para int, fileContent []string, filePath string) error {
	err := util.String2tree(fileContent)
	if err != nil {
		return err
	}
	if !util.IsInit() {
		return errors.New("OutputAsFile(): No file in workspace")
	}

	tree := util.GetRoot()
	if para != 0 {
		// 向文件输入
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return errors.New(err.Error())
		}

		// 清空文件内容
		err = file.Truncate(0)
		if err != nil {
			return errors.New(err.Error())
		}

		for _, child := range tree.Children {
			recurOutputAsFile(child, file)
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	} else {
		// 向终端输入
		for _, child := range tree.Children {
			recurOutputAsFile(child, nil)
		}
	}
	return nil
}

func recurOutputAsTree(prefix string, treeOut TreeOut) error {
	if treeOut == nil {
		return nil
	}
	l := len(treeOut.GetChildren())
	for i, child := range treeOut.GetChildren() {
		fmt.Print(prefix)
		var err error
		if i != l-1 {
			fmt.Print(branch)
			fmt.Println(child.Name())
			err = recurOutputAsTree(prefix+notSpace, child)
		} else {
			fmt.Print(endBranch)
			fmt.Println(child.Name())
			err = recurOutputAsTree(prefix+space, child)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func OutputAsTree(fileContent []string) error {
	err := util.String2tree(fileContent)
	if err != nil {
		return err
	}
	if !util.IsInit() {
		return errors.New("OutputAsTree(): No file in workspace")
	}

	tree := util.GetRoot()
	//for _, child := range tree.children {
	err = recurOutputAsTree("", tree)
	if err != nil {
		return err
	}
	//}
	return nil
}

func OutputAsDir(content string, fileContent []string) error {
	err := util.String2tree(fileContent)
	if err != nil {
		return err
	}
	if !util.IsInit() {
		return errors.New("OutputAsDir(): No file in workspace")
	}

	_, node := util.GetNodeByContent(content)
	if node == nil {
		return errors.New("OutputAsDir(): No such node")
	}
	fmt.Print(endBranch)
	fmt.Println(node.Content)
	//for _, child := range node.children {
	err = recurOutputAsTree(space, node)
	if err != nil {
		return err
	}
	//}
	return nil
}

// 对每一个节点 首先判断他是不是最后一个
// 如果是最后一个，那么就输出└──
// 如果不是最后一个，那么就输出├──
// 然后输出他的内容
// 然后输出他的子节点
// └── 我的资源
// ├── 程序设计
// │ └── 软件设计
// │ └── 设计模式
// │ ├── 1. 观察者模式
// │ ├── 2. 策略模式
// │ └── 3. 组合模式
// └── ⼯具箱
// └── Adobe
