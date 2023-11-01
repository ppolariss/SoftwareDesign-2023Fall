package main

import (
	// "bufio"
	// "container/list"
	// e "design/myError"
	// "fmt"
	// "os"
	// "strings"
	"design/command"
	// "fmt"
	// "design/tree"
)

func main() {
	// var a tree.node
	for {
		err := command.Do()
		if err != nil {
			// fmt.Println(err)
			break
		}
	}
}

// 定义宏 定义函数指针
// 将输入 映射到结构体

// // 接收者
// type Receiver struct{}

// func (r *Receiver) Action1() {
// 	fmt.Println("Receiver: Action1 executed")
// }

// func (r *Receiver) Action2() {
// 	fmt.Println("Receiver: Action2 executed")
// }

// // 调用者
// type Invoker struct {
// 	commands []Command
// }

// func (i *Invoker) AddCommand(command Command) {
// 	i.commands = append(i.commands, command)
// }

// func (i *Invoker) ExecuteCommands() {
// 	for _, command := range i.commands {
// 		command.Execute()
// 	}
// }

// func main() {
// 	// for{}
// 	c, err := readCommand()
// 	if c == nil {
// 		// if(str!="exit")
// 		fmt.Println(err)
// 		// panic("")
// 		return
// 	}
// 	err = c.Execute()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// 创建接收者对象
// 	receiver := &Receiver{}

// 	// 创建具体命令对象，并关联接收者
// 	command1 := &ConcreteCommand1{receiver: receiver}
// 	command2 := &ConcreteCommand2{receiver: receiver}

// 	// 创建调用者对象，并添加命令
// 	invoker := &Invoker{}
// 	invoker.AddCommand(command1)
// 	invoker.AddCommand(command2)

// 	// 执行命令
// 	invoker.ExecuteCommands()
// }
