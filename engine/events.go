package engine

import "cli-graphics/utils"

type Event interface{}

type KeyEvent struct {
	Key string
}

type TerminalResizeEvent struct {
	Width, Height int
}

type MouseEvent struct {
	X, Y        int
	MouseButton utils.MouseBtn
	MouseAction utils.MouseAction
}
