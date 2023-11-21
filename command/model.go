package command

import (
	"design/commandManager"
	. "design/interfaces"
	"design/log"
	"design/workspace"
)

var curWorkspace = &workspace.Workspace{}

var commandsMapper = map[string]Command{
	"load":        &load{},
	"save":        &save{},
	"insert":      &insert{},
	"delete":      &deleteCommand{},
	"append-head": &appendHead{},
	"append-tail": &appendTail{},
	"undo":        &commandManager.Undo{},
	"redo":        &commandManager.Redo{},
	"list":        &list{},
	"list-tree":   &listTree{},
	"dir-tree":    &dirTree{},
	"history":     &log.History{},
	"stats":       &stats{},
	"ls":          &ls{},
}
