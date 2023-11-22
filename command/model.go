package command

import (
	"design/commandManager"
	"design/editor"
	. "design/interfaces"
	"design/log"
	"design/output"
	"design/workspace"
)

var commandsMapper = map[string]Command{
	"insert":         &editor.InsertCommand{},
	"delete":         &editor.DeleteCommand{},
	"append-head":    &editor.AppendHead{},
	"append-tail":    &editor.AppendTail{},
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
