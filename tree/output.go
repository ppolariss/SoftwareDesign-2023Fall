package tree

import (
	e "design/myError"
	"design/util"
	"fmt"
	"os"
)

const endBranch = "└── "
const branch = "├── "
const space = "    "
const notSpace = "│   "

func dumpNode(node *Node) {
	fmt.Println(node.content)
	if len(node.children) != 0 {
		fmt.Println("children: ")
	}
	for _, child := range node.children {
		dumpNode(child)
	}
}

func Dump() {
	fmt.Println("Dump")
	root := GetRoot()
	for _, child := range root.children {
		dumpNode(child)
	}
}

func recurOutputAsFile(node *Node, file *os.File) {
	if node.grade != 0 {
		for i := 0; i < node.grade; i++ {
			util.Output("#", file)
			
		}
		util.Output(" ", file)
		
	}
	
	util.Output(node.content+"\n", file)
	for _, child := range node.children {
		recurOutputAsFile(child, file)
	}
}

// para: 0: output to terminal; 1: output to file
func OutputAsFile(para int) error {
	if !IsInit() {
		return e.NewMyError("OutputAsFile(): No file in workspace")
	}

	tree := GetRoot()
	if para != 0 {
		// 向文件输入
		file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return e.NewMyError(err.Error())
		}

		// 清空文件内容
		err = file.Truncate(0)
		if err != nil {
			return e.NewMyError(err.Error())
		}

		for _, child := range tree.children {
			recurOutputAsFile(child, file)
		}
		defer file.Close()
	} else {
		// 向终端输入
		for _, child := range tree.children {
			recurOutputAsFile(child, nil)
		}
	}
	return nil
}

func recurOutputAsTree(prefix string, node *Node) error {
	fmt.Print(prefix)
	rank, err := getRankofParent(node)
	if err != nil {
		return err
	} else if rank == -1 {
		fmt.Print(endBranch)
		prefix += space
	} else {
		fmt.Print(branch)
		prefix += notSpace
	}
	fmt.Println(node.content)

	for _, child := range node.children {
		err = recurOutputAsTree(prefix, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func OutputAsTree() error {
	if !IsInit() {
		return e.NewMyError("OutputAsTree(): No file in workspace")
	}

	tree := GetRoot()
	for _, child := range tree.children {
		err := recurOutputAsTree("", child)
		if err != nil {
			return err
		}
	}
	return nil
}


func OutputAsDir(content string) error {
	if !IsInit() {
		return e.NewMyError("OutputAsDir(): No file in workspace")
	}

	_, node := getNodeByContent(content)
	if node == nil {
		return e.NewMyError("OutputAsDir(): No such node")
	}
	fmt.Print(endBranch)
	fmt.Println(node.content)
	for _, child := range node.children {
		err := recurOutputAsTree(space, child)
		if err != nil {
			return err
		}
	}
	return nil
}

// 对每一个节点 首先判断他是不是最后一个
// 如果是最后一个，那么就输出└──
// 如果不是最后一个，那么就输出├──
// 然后输出他的内容
// 然后输出他的子节点

// func OutputAsTree() error {
// 	if !IsInit() {
// 		return e.NewMyError("OutputAsTree(): No file in workspace")
// 	}

// 	tree := GetRoot()
// 	for _, child := range tree.children {
// 		// recurOutputAsTree(child, 0)

// 	}
// 	return nil
// }
// └── 我的资源
// ├── 程序设计
// │ └── 软件设计
// │ └── 设计模式
// │ ├── 1. 观察者模式
// │ ├── 2. 策略模式
// │ └── 3. 组合模式
// └── ⼯具箱
// └── Adobe
