package output

import (
	. "design/interfaces"
	"design/util"
	"errors"
	"fmt"
)

const endBranch = "└── "
const branch = "├── "
const space = "    "
const notSpace = "│   "

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

func AsTree(fileContent []string) error {
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

func AsDir(content string, fileContent []string) error {
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
