package engine

type Focusable interface {
	Component

	SetFocus(focused bool)
	isFocused() bool

	HandleKey(key string) bool
}
