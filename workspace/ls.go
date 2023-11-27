package workspace

import (
	"design/output"
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
	updateWorkspace(CurWorkspace)
	return output.Ls(Path, func(name string) string {
		for _, n := range AllWorkspaces {
			if n.FileName == name {
				return star
			}
		}
		return ""
	})
}

func (c *Ls) CallSelf() string {
	return "ls"
}
