package command

import (
	"design/commandManager"
	. "design/interfaces"
	"design/log"
	"design/output"
	"design/workspace"
)

var commandsMapper = map[string]Command{
	"insert":         &InsertCommand{},
	"delete":         &DeleteCommand{},
	"append-head":    &AppendHead{},
	"append-tail":    &AppendTail{},
	"list":           &List{},
	"list-tree":      &ListTree{},
	"dir-tree":       &DirTree{},
	"ls":             &output.Ls{},
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
}
