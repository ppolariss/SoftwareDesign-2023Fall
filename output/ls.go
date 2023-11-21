package output

import (
	"errors"
)

type Ls struct {
}

func (c *Ls) SetArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("ls: args error")
	}
	return nil
}

func (c *Ls) Execute() error {
	return recurOutputAsTree("", &File{path: "./files/"})
}
func (c *Ls) CallSelf() string {
	return "ls"
}
