package engine

type Container interface {
	Component

	AddChild(child Component)
	FindNextFocusableChild(st int) int
	Flush()
}
