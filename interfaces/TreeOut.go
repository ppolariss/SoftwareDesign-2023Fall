package interfaces

type TreeOut interface {
	Name() string
	Children() []TreeOut
}
