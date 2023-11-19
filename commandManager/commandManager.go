package commandManager

import (
	. "design/interfaces"
)

// record commands which has reverse_command except Undo
// record in undoHistory rather than go to canUnDoHistory when Undo
// flush undoHistory when log
var canUnDoHistory []UndoableCommand

var canUnDoPointer int

// pointer: - 1 ~ len - 1
// means next Undo Command
func next() UndoableCommand {
	if canUnDoPointer >= 0 && canUnDoPointer < len(canUnDoHistory) {
		canUnDoPointer--
		return canUnDoHistory[canUnDoPointer+1]
	} else {
		return nil
	}
}

func previous() UndoableCommand {
	if canUnDoPointer >= -1 && canUnDoPointer < len(canUnDoHistory)-1 {
		canUnDoPointer++
		return canUnDoHistory[canUnDoPointer]
	} else {
		return nil
	}
}
