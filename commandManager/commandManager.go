package commandManager

import (
	. "design/interfaces"
	. "design/workspace"
)

// record commands which has reverse_command except Undo
// record in undoHistory rather than go to canUnDoHistory when Undo
// flush undoHistory when log
//var canUnDoHistory []UndoableCommand
//
//var canUnDoPointer int

// pointer: - 1 ~ len - 1
// means next Undo Command
func next() UndoableCommand {
	if CurWorkspace.UndoableCommandPointer >= 0 && CurWorkspace.UndoableCommandPointer < len(CurWorkspace.UndoableCommandHistory) {
		CurWorkspace.UndoableCommandPointer--
		updateDirty(CurWorkspace.UndoableCommandPointer)
		return CurWorkspace.UndoableCommandHistory[CurWorkspace.UndoableCommandPointer+1]
	} else {
		return nil
	}
}

func previous() UndoableCommand {
	if CurWorkspace.UndoableCommandPointer >= -1 && CurWorkspace.UndoableCommandPointer < len(CurWorkspace.UndoableCommandHistory)-1 {
		CurWorkspace.UndoableCommandPointer++
		updateDirty(1)
		return CurWorkspace.UndoableCommandHistory[CurWorkspace.UndoableCommandPointer]
	} else {
		return nil
	}
}

func updateDirty(p int) {
	if p == -1 {
		CurWorkspace.Dirty = false
	} else {
		CurWorkspace.Dirty = true
	}
}
