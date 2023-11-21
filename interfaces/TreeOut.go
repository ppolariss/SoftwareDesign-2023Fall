package interfaces

type TreeOut interface {
	Name() string
	GetChildren() []TreeOut
}
