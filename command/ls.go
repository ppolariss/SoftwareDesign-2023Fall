package command

import e "design/myError"

type ls struct {
}

func (c *ls) SetArgs(args []string) error {
	if len(args) != 1 {
		return e.NewMyError("ls: args error")
	}
	return nil
}

func (c *ls) Execute() error {
	// TODO
	return nil
}
func (c *ls) CallSelf() string {
	return "ls"
}
