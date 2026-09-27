package engine

type Event interface{}

type KeyEvent struct {
	Key string
}

type TerminalResizeEvent struct {
	Width, Height int
}
