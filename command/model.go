package command

import (
	"design/commandManager"
	. "design/interfaces"
	"design/log"
	"design/workspace"
)

type getCommand func() Command

var commandsMapper = map[string]getCommand{
	"insert":         func() Command { return &InsertCommand{} },
	"delete":         func() Command { return &DeleteCommand{} },
	"append-head":    func() Command { return &AppendHead{} },
	"append-tail":    func() Command { return &AppendTail{} },
	"list":           func() Command { return &List{} },
	"list-tree":      func() Command { return &ListTree{} },
	"dir-tree":       func() Command { return &DirTree{} },
	"ls":             func() Command { return &workspace.Ls{} },
	"load":           func() Command { return &workspace.Load{} },
	"save":           func() Command { return &workspace.Save{} },
	"exit":           func() Command { return &workspace.Exit{} },
	"close":          func() Command { return &workspace.Close{} },
	"list-workspace": func() Command { return &workspace.List{} },
	"open":           func() Command { return &workspace.Open{} },
	"undo":           func() Command { return &commandManager.Undo{} },
	"redo":           func() Command { return &commandManager.Redo{} },
	"history":        func() Command { return &log.History{} },
	"stats":          func() Command { return &log.Stats{} },
}

//var commandsMapper = map[string]Command{
//	"insert":         &InsertCommand{},
//	"delete":         &DeleteCommand{},
//	"append-head":    &AppendHead{},
//	"append-tail":    &AppendTail{},
//	"list":           &List{},
//	"list-tree":      &ListTree{},
//	"dir-tree":       &DirTree{},
//	"ls":             &output.Ls{},
//	"load":           &workspace.Load{},
//	"save":           &workspace.Save{},
//	"exit":           &workspace.Exit{},
//	"close":          &workspace.Close{},
//	"list-workspace": &workspace.List{},
//	"open":           &workspace.Open{},
//	"undo":           &commandManager.Undo{},
//	"redo":           &commandManager.Redo{},
//	"history":        &log.History{},
//	"stats":          &log.Stats{},
//}
