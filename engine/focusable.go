package engine

type Focusable interface {
	Component

	SetFocus(focused bool)
	IsFocused() bool

	HandleKey(key string) bool
}
