package command

import (
	"fmt"
	e "design/myError"
	"design/tree"
)

// // 具体命令2
type save struct {
	// receiver *Receiver
}

func (c *save) Execute() error {

	fmt.Println("save")
	return tree.OutputAsFile(1)
	// c.receiver.Action2()
}

func (c *save) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("save: args error")
		// return "save: args error"
	}
	return nil
}