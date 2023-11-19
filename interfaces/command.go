package interfaces

type Command interface {
	SetArgs([]string) error
	Execute() error
	CallSelf() string
}

type UndoableCommand interface {
	Command
	UndoExecute() error
}
