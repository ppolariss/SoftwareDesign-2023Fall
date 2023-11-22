package command

import (
	"design/commandManager"
	. "design/interfaces"
	"design/log"
	"design/output"
	"design/workspace"
)

var commandsMapper = map[string]Command{
	"insert":         &insert{},
	"delete":         &deleteCommand{},
	"append-head":    &appendHead{},
	"append-tail":    &appendTail{},
	"load":           &workspace.Load{},
	"save":           &workspace.Save{},
	"exit":           &workspace.Exit{},
	"close":          &workspace.Close{},
	"list-workspace": &workspace.List{},
	"open":           &workspace.Open{},
	"undo":           &commandManager.Undo{},
	"redo":           &commandManager.Redo{},
	"history":        &log.History{},
	"stats":          &log.Stats{},
	"list":           &output.List{},
	"list-tree":      &output.ListTree{},
	"dir-tree":       &output.DirTree{},
	"ls":             &output.Ls{},
}
